package main

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/goccy/go-yaml"
	management "github.com/jasonbot/windows-msix-handover-app/management"
	"github.com/shirou/gopsutil/process"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-winrtapi/winrt"
	"golang.org/x/sys/windows/registry"
)

type YamlUpdateFile struct {
	Url    string `yaml:"url"`
	Sha512 string `yaml:"sha512"`
	Size   uint64 `yaml:"size"`
}

type YamlUpdateStruct struct {
	Files       []YamlUpdateFile `yaml:"files"`
	ReleaseDate string           `yaml:"releaseDate"`
	Version     string           `yaml:"version"`
	Sha512      string           `yaml:"sha512"`
	Path        string           `yaml:"path"`
}

type CPUArchitecture string

const (
	ArchArm64 CPUArchitecture = "arm64"
	ArchAmd64 CPUArchitecture = "amd64"
)

type DesktopProduct struct {
	ProductName  string
	Architecture CPUArchitecture
}

var DesktopProductFeeds map[DesktopProduct]string

func init() {
	DesktopProductFeeds = map[DesktopProduct]string{
		DesktopProduct{ProductName: "Notion", Architecture: ArchArm64}:     "https://desktop-release.notion-static.com/arm64-msix.yml",
		DesktopProduct{ProductName: "Notion Dev", Architecture: ArchArm64}: "https://dev-desktop-release.s3.us-west-2.amazonaws.com/arm64-msix.yml",
		DesktopProduct{ProductName: "Notion Stg", Architecture: ArchArm64}: "https://stg-desktop-release.s3.us-west-2.amazonaws.com/arm64-msix.yml",
		DesktopProduct{ProductName: "Notion", Architecture: ArchAmd64}:     "https://desktop-release.notion-static.com/msix.yml",
		DesktopProduct{ProductName: "Notion Dev", Architecture: ArchAmd64}: "https://dev-desktop-release.s3.us-west-2.amazonaws.com/msix.yml",
		DesktopProduct{ProductName: "Notion Stg", Architecture: ArchAmd64}: "https://stg-desktop-release.s3.us-west-2.amazonaws.com/msix.yml",
	}
}

func stopAppIfRunning(appName string) {
	exeName := fmt.Sprintf("%v.exe", appName)

	findRootProcess := func(p *process.Process) *process.Process {
		exeName, _ := p.Exe()

		for {
			if parentProc, err := p.Parent(); parentProc != nil && err == nil {
				if parentExeName, _ := parentProc.Exe(); parentExeName != exeName {
					return p
				}
				p = parentProc
			} else {
				fmt.Println()
				return parentProc
			}
		}
	}

	if p, err := process.Processes(); err == nil {
		for _, runningProcess := range p {
			if e, err := runningProcess.Exe(); err == nil {
				pb := filepath.Base(filepath.Clean(e))
				if pb == exeName {
					proc := findRootProcess(runningProcess)
					fmt.Println("Found process", pb, proc.Pid)
					proc.Terminate()
					for pr, _ := proc.IsRunning(); pr; pr, _ = proc.IsRunning() {
						fmt.Println("Waiting to quit")
					}
				}
			}
		}
	}
}

func uninstallWin32AppIfInstalled(appName string) {
	ui := `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`

	for _, tk := range []registry.Key{registry.CURRENT_USER, registry.LOCAL_MACHINE} {
		if k, err := registry.OpenKey(tk, ui, registry.ALL_ACCESS); err == nil {
			defer k.Close()

			if subkeys, err := k.ReadSubKeyNames(2 << 16); subkeys != nil {
				for _, subkey := range subkeys {
					subkeyPath := fmt.Sprintf(`%v\%v`, ui, subkey)
					if sk, err := registry.OpenKey(registry.CURRENT_USER, subkeyPath, registry.ALL_ACCESS); err == nil {
						defer sk.Close()
						displayVersion, _, _ := sk.GetStringValue("DisplayVersion")
						displayName, _, _ := sk.GetStringValue("DisplayName")
						quietUninstallString, _, _ := sk.GetStringValue("QuietUninstallString")
						if quietUninstallString == "" {
							quietUninstallString, _, _ = sk.GetStringValue("UninstallString")
						}
						publisher, _, _ := sk.GetStringValue("Publisher")

						desiredProductName := fmt.Sprintf("%v %v", appName, displayVersion)

						if displayName == desiredProductName {
							fmt.Println("Found it:", quietUninstallString, publisher)
							return
						}
					}
				}
			} else {
				fmt.Println("ERROR", err)
			}
		}
	}
}

func findLatestMSIXUpdate(channelYamlURL string) (string, int64, string, error) {
	req, _ := http.NewRequest("GET", channelYamlURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		return "", 0, "", fmt.Errorf("fetching %v returned %v", channelYamlURL, resp.StatusCode)
	}
	defer resp.Body.Close()

	var config YamlUpdateStruct
	if yamlSource, err := io.ReadAll(resp.Body); err == nil {
		yaml.Unmarshal(yamlSource, &config)
	} else {
		return "", 0, "", err
	}

	log.Println("CXonfig", config)

	for _, file := range config.Files {
		if strings.HasSuffix(strings.ToLower(file.Url), ".msix") {
			installerUrl := strings.Replace(file.Url, "-x64", "", 1)
			if computedUrl, err := url.JoinPath(
				channelYamlURL,
				"..",
				strings.Replace(
					file.Url,
					"-x64",
					"",
					1,
				),
			); err == nil {
				installerUrl = computedUrl
			}

			log.Println("URL:", config.Path, "->", installerUrl)

			return installerUrl, int64(file.Size), file.Sha512, nil
		}
	}

	return "", 0, "", errors.New("no update msix found")
}

type writerWrapper struct {
	hasher     hash.Hash
	Out        io.Writer
	bytesSoFar int64
	TotalBytes int64
}

func (w *writerWrapper) Write(p []byte) (n int, err error) {
	if w.hasher == nil {
		w.hasher = sha512.New()
	}
	w.hasher.Write(p)
	w.bytesSoFar += int64(len(p))
	return w.Out.Write(p)
}

func (w *writerWrapper) Hash() string {
	if w.hasher != nil {
		bs := w.hasher.Sum(nil)
		return base64.StdEncoding.EncodeToString(bs)
	}
	return ""
}

func downloadMSIXToDownloadsFolder(msixURL string, fileSize int64, sha512 string) string {
	u, _ := url.Parse(msixURL)
	if u == nil {
		return ""
	}

	homeDir, _ := os.UserHomeDir()
	downloadsPath := filepath.Join(homeDir, "Downloads")

	pathParts := strings.Split(u.Path, "/")
	fileName := pathParts[len(pathParts)-1]
	ext := path.Ext(fileName)
	fileBase := path.Join(downloadsPath, fileName[:len(fileName)-len(ext)])

	num := 2
	_, err := os.Stat(fileName)
	for !errors.Is(err, os.ErrNotExist) {
		fileName = fmt.Sprintf("%v (%v)%v", fileBase, num, ext)
		_, err = os.Stat(fileName)
		num += 1
	}

	req, _ := http.NewRequest("GET", msixURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		return ""
	}
	defer resp.Body.Close()

	f, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

	writer := writerWrapper{
		Out:        f,
		TotalBytes: fileSize,
	}
	io.Copy(&writer, resp.Body)

	if writer.Hash() != sha512 {
		return ""
	}

	return fileName
}

func Co_Await[R comparable](operation *winrt.IAsyncOperation[R],
	onComplete func(*winrt.IAsyncOperation[R], winrt.AsyncStatus) error,
) {
	done := make(chan struct{})

	operation.Put_Completed(
		func(operation *winrt.IAsyncOperation[R], asyncStatus winrt.AsyncStatus) com.Error {
			returnVal := com.OK
			if err := onComplete(operation, asyncStatus); err != nil {
				returnVal = com.FAIL
			}
			//win32.PostThreadMessage(com.GetContext().TID, win32.WM_QUIT, 0, 0)
			done <- struct{}{}

			return returnVal
		},
	)

	//com.MessageLoop()
	<-done
}

func Co_AwaitWithProgress[R, P comparable](operation *winrt.IAsyncOperationWithProgress[R, P],
	onProgress func(*winrt.IAsyncOperationWithProgress[R, P], P) error,
	onComplete func(*winrt.IAsyncOperationWithProgress[R, P], winrt.AsyncStatus) error,
) {
	done := make(chan struct{})

	operation.Put_Progress(
		func(operation *winrt.IAsyncOperationWithProgress[R, P], progress P) com.Error {
			if onProgress != nil {
				if err := onProgress(operation, progress); err != nil {
					return com.FAIL
				}
			}
			return com.OK
		},
	)
	operation.Put_Completed(
		func(operation *winrt.IAsyncOperationWithProgress[R, P], asyncStatus winrt.AsyncStatus) com.Error {
			returnVal := com.OK
			if err := onComplete(operation, asyncStatus); err != nil {
				returnVal = com.FAIL
			}
			// win32.PostThreadMessage(com.GetContext().TID, win32.WM_QUIT, 0, 0)
			done <- struct{}{}

			return returnVal
		},
	)

	// com.MessageLoop()
	<-done
}

func installMSIXFromDownloadsFolder(msixPath string) {
	runtime.LockOSThread()
	winrt.Initialize()
	// defer winrt.Uninitialize()

	msixURI := fmt.Sprint("file://", strings.ReplaceAll(msixPath, "\\", "/"))
	log.Println("MSIX", msixURI)
	uri := winrt.NewUri_CreateUri(msixURI)
	log.Println("MSIX URL:", msixURI)

	pm := management.NewPackageManager()
	if pm.IUnknown.GetIUnknown() == nil {
		log.Println("No package manager")
		return
	}

	op := pm.AddPackageAsync(
		uri.IUriRuntimeClass,
		nil,
		management.DeploymentOptions_ForceTargetApplicationShutdown,
	)

	Co_AwaitWithProgress(
		op,
		func(
			_ *winrt.IAsyncOperationWithProgress[*management.IDeploymentResult, management.DeploymentProgress],
			progress management.DeploymentProgress,
		) error {
			log.Println("Percentage:", progress.Percentage)
			return nil
		},
		func(
			asyncInfo *winrt.IAsyncOperationWithProgress[*management.IDeploymentResult, management.DeploymentProgress],
			asyncStatus winrt.AsyncStatus,
		) error {
			log.Println("Finished:", asyncStatus)

			if asyncStatus == winrt.AsyncStatus_Completed || asyncStatus == winrt.AsyncStatus_Error {
				if r := asyncInfo.GetResults(); r != nil {
					if asyncStatus == winrt.AsyncStatus_Error {
						errorText := r.Get_ErrorText()
						log.Println("Error:", errorText)
					} else {
						log.Println("It worked")
					}
				}
			}
			return nil
		},
	)
}

func Co_CreateInstanceByClassID[T comparable](clsid string, iid syscall.GUID) (*T, error) {
	var p *T
	hs := winrt.NewHStr(clsid)
	hr := win32.RoGetActivationFactory(hs.Ptr, &iid, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return p, fmt.Errorf("error in processing: %v", win32.HRESULT_ToString(hr))
	}
	com.AddToScope(p)
	return p, nil
}

func runInstalledApp() {
	runtime.LockOSThread()
	winrt.Initialize()
	// defer winrt.Uninitialize()

	pm := management.NewPackageManager()
	if pm.IUnknown.GetIUnknown() == nil {
		log.Println("No package manager")
		return
	}

	if packages, err := pm.FindPackages(); err == nil {
		iterator := packages.First()

		for iterator.Get_HasCurrent() == true {
			currentPackage := iterator.Get_Current()
			if currentPackage != nil {
				il := currentPackage.Get_Id()
				msixURI := fmt.Sprint("shell:AppsFolder\\", url.PathEscape(il.Get_FullName()))
				uri := winrt.NewUri_CreateUri(msixURI)
				log.Println("App", il.Get_Name(), msixURI)

				if il.Get_Name() == "MPOSSIBL:E" {
					ls, _ := Co_CreateInstanceByClassID[winrt.ILauncherStatics](
						"Windows.System.Launcher",
						winrt.IID_ILauncherStatics,
					)

					proc := ls.LaunchUriAsync(uri.IUriRuntimeClass)
					Co_Await[bool](proc, func(*winrt.IAsyncOperation[bool], winrt.AsyncStatus) error {
						return nil
					})

					return
				}
			}
			iterator.MoveNext()
		}
	}
}

func main() {
	app := DesktopProduct{
		ProductName:  "Notion Dev",
		Architecture: CPUArchitecture(runtime.GOARCH),
	}

	stopAppIfRunning(app.ProductName)

	msixUrl, fileSize, sha512, err := findLatestMSIXUpdate(DesktopProductFeeds[app])
	if err == nil {
		msixPath := downloadMSIXToDownloadsFolder(msixUrl, fileSize, sha512)

		if msixPath != "" {
			installMSIXFromDownloadsFolder(msixPath)
			uninstallWin32AppIfInstalled(app.ProductName)
		}
	}
}
