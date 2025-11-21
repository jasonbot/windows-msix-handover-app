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

	"github.com/goccy/go-yaml"
	management "github.com/jasonbot/windows-msix-handover-app/management"
	"github.com/shirou/gopsutil/process"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-winrtapi/winrt"
	"golang.org/x/sys/windows/registry"
)

type updateFile struct {
	Url    string `yaml:"url"`
	Sha512 string `yaml:"sha512"`
	Size   uint64 `yaml:"size"`
}

type updateStruct struct {
	Files       []updateFile `yaml:"files"`
	ReleaseDate string       `yaml:"releaseDate"`
	Version     string       `yaml:"version"`
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
					return
				}
			}
		}
	}
}

func uninstallAppIfInstalled(appName string) {
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

	var config updateFile
	if yamlSource, err := io.ReadAll(resp.Body); err == nil {
		yaml.Unmarshal(yamlSource, &config)
	} else {
		return "", 0, "", err
	}

	return config.Url, int64(config.Size), config.Sha512, nil
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

func Co_Await[R, P comparable](operation *winrt.IAsyncOperation[*R],
	onComplete func(*winrt.IAsyncOperation[*R], winrt.AsyncStatus) error,
) {
	done := make(chan struct{})

	operation.Put_Completed(
		func(operation *winrt.IAsyncOperation[*R], asyncStatus winrt.AsyncStatus) com.Error {
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

func Co_AwaitWithProgress[R, P comparable](operation *winrt.IAsyncOperationWithProgress[*R, P],
	onProgress func(*winrt.IAsyncOperationWithProgress[*R, P], P) error,
	onComplete func(*winrt.IAsyncOperationWithProgress[*R, P], winrt.AsyncStatus) error,
) {
	done := make(chan struct{})

	operation.Put_Progress(
		func(operation *winrt.IAsyncOperationWithProgress[*R, P], progress P) com.Error {
			if onProgress != nil {
				if err := onProgress(operation, progress); err != nil {
					return com.FAIL
				}
			}
			return com.OK
		},
	)
	operation.Put_Completed(
		func(operation *winrt.IAsyncOperationWithProgress[*R, P], asyncStatus winrt.AsyncStatus) com.Error {
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

func installMSIXFromDownloadsFolder(msixPath string) {
	runtime.LockOSThread()
	winrt.Initialize()
	defer winrt.Uninitialize()

	msixURI := fmt.Sprint("file://", url.PathEscape(msixPath))
	uri := winrt.NewUri_CreateUri(msixURI)

	pm := management.NewPackageManager()
	if pm.IUnknown.GetIUnknown() == nil {
		log.Println("Ooof")
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

type appConfig struct {
	AppName        string
	ChannelYamlURL string
}

func main() {
	app := appConfig{
		AppName:        "Notion Dev",
		ChannelYamlURL: "https://fart.com/",
	}

	stopAppIfRunning(app.AppName)
	uninstallAppIfInstalled(app.AppName)

	msixUrl, fileSize, sha512, err := findLatestMSIXUpdate(app.ChannelYamlURL)
	if err != nil {
		msixPath := downloadMSIXToDownloadsFolder(msixUrl, fileSize, sha512)
		installMSIXFromDownloadsFolder(msixPath)
	}
}
