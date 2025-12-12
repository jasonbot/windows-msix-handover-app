package main

import (
	"bufio"
	"crypto/sha512"
	"debug/pe"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/goccy/go-yaml"
	"github.com/jasonbot/windows-msix-handover-app/channels"
	"github.com/jasonbot/windows-msix-handover-app/checklist"
	"github.com/jasonbot/windows-msix-handover-app/config"
	management "github.com/jasonbot/windows-msix-handover-app/management"
	"github.com/shirou/gopsutil/process"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-winrtapi/winrt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func IsArm64() bool {
	if runtime.GOARCH == "arm64" {
		return true
	}

	// Detect if we're running
	var processMachine, nativeMachine uint16
	err := windows.IsWow64Process2(windows.CurrentProcess(), &processMachine, &nativeMachine)
	if err == nil {
		return nativeMachine == pe.IMAGE_FILE_MACHINE_ARM64
	}

	return false
}

func Ro_Await[R comparable](operation *winrt.IAsyncOperation[R],
	onComplete func(*winrt.IAsyncOperation[R], winrt.AsyncStatus) error,
) {
	done := make(chan struct{})

	operation.Put_Completed(
		func(operation *winrt.IAsyncOperation[R], asyncStatus winrt.AsyncStatus) com.Error {
			returnVal := com.OK
			if err := onComplete(operation, asyncStatus); err != nil {
				returnVal = com.FAIL
			}
			done <- struct{}{}

			return returnVal
		},
	)

	<-done
}

func Ro_AwaitWithProgress[R, P comparable](operation *winrt.IAsyncOperationWithProgress[R, P],
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
			done <- struct{}{}

			return returnVal
		},
	)

	<-done
}

func Ro_CreateInstanceByClassID[T comparable](clsid string, iid syscall.GUID) (*T, error) {
	var p *T
	hs := winrt.NewHStr(clsid)
	hr := win32.RoGetActivationFactory(hs.Ptr, &iid, unsafe.Pointer(&p))
	if win32.FAILED(hr) {
		return p, fmt.Errorf("error in processing: %v", win32.HRESULT_ToString(hr))
	}
	com.AddToScope(p)
	return p, nil
}

func RunElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		log.Println(err)
	}
}

func IsRunningElevated() bool {
	return windows.GetCurrentProcessToken().IsElevated()
}

func stopAppIfRunning(appName string, rs checklist.RunStep) {
	exeName := fmt.Sprintf("%v.exe", appName)
	rs.SetState(checklist.StepInProgress)
	rs.SetMessage("Looking for " + exeName)

	findRootProcess := func(p *process.Process) *process.Process {
		exeName, _ := p.Exe()

		for {
			if parentProc, err := p.Parent(); parentProc != nil && err == nil {
				if parentExeName, _ := parentProc.Exe(); parentExeName != exeName {
					return p
				}
				p = parentProc
			} else {
				log.Println()
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
					log.Println("Found process", pb, proc.Pid)
					proc.Terminate()
					for pr, _ := proc.IsRunning(); pr; pr, _ = proc.IsRunning() {
						rs.SetMessage(fmt.Sprintf("Witing for PID %v to exit", proc.Pid))
						log.Println("Waiting to quit")
					}
				}
			}
		}
	}
	rs.SetState(checklist.StepSuccess)
}

func uninstallWin32AppIfInstalled(appName string, rs checklist.RunStep) {
	rs.SetState(checklist.StepInProgress)
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
							log.Println("Found it:", quietUninstallString, publisher)
							log.Println("When done:", exec.Command(quietUninstallString).Run())
							rs.SetState(checklist.StepSuccess)
							return
						}
					}
				}
			} else {
				rs.SetState(checklist.StepError)
				log.Println("Error reading keys", err)
			}
		}
	}
}

func findLatestMSIXUpdate(channelYamlURL string, rs checklist.RunStep) (string, int64, string, error) {
	rs.SetState(checklist.StepInProgress)
	req, _ := http.NewRequest("GET", channelYamlURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		rs.SetMessage(fmt.Sprintf("fetching %v returned %v", channelYamlURL, resp.StatusCode))
		rs.SetState(checklist.StepError)
		return "", 0, "", fmt.Errorf("fetching %v returned %v", channelYamlURL, resp.StatusCode)
	}
	defer resp.Body.Close()

	var config channels.YamlUpdateStruct
	if yamlSource, err := io.ReadAll(resp.Body); err == nil {
		yaml.Unmarshal(yamlSource, &config)
	} else {
		return "", 0, "", err
	}

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
	Progress   func(int64, int64)
}

func (w *writerWrapper) Write(p []byte) (n int, err error) {
	if w.hasher == nil {
		w.hasher = sha512.New()
	}
	w.hasher.Write(p)
	w.bytesSoFar += int64(len(p))
	log.Println("Progress:", w.bytesSoFar, "/", w.TotalBytes)
	if w.Progress != nil {
		w.Progress(w.bytesSoFar, w.TotalBytes)
	}
	return w.Out.Write(p)
}

func (w *writerWrapper) Hash() string {
	if w.hasher != nil {
		bs := w.hasher.Sum(nil)
		return base64.StdEncoding.EncodeToString(bs)
	}
	return ""
}

func shaForPath(filePath string) string {
	hasher := sha512.New()
	if s, err := os.ReadFile(filePath); err == nil {
		hasher.Write(s)
		bs := hasher.Sum(nil)
		return base64.StdEncoding.EncodeToString(bs)
	}
	return "no"
}

func downloadMSIXToDownloadsFolder(msixURL string, fileSize int64, expectedSha512Sum string, rs checklist.RunStep) (string, bool) {
	rs.SetState(checklist.StepInProgress)
	u, _ := url.Parse(msixURL)
	if u == nil {
		return "", false
	}

	homeDir, _ := os.UserHomeDir()
	downloadsPath := filepath.Join(homeDir, "Downloads")

	pathParts := strings.Split(u.Path, "/")
	fileName := pathParts[len(pathParts)-1]
	ext := path.Ext(fileName)
	fileBase := path.Join(downloadsPath, fileName[:len(fileName)-len(ext)])

	fileName = path.Join(downloadsPath, fileName)
	num := 2
	_, err := os.Stat(fileName)
	for !errors.Is(err, os.ErrNotExist) {
		if shaForPath(fileName) == expectedSha512Sum {
			log.Println("Found already downloaded", fileName)
			return fileName, false
		}

		fileName = fmt.Sprintf("%v (%v)%v", fileBase, num, ext)
		_, err = os.Stat(fileName)
		num += 1
	}

	req, _ := http.NewRequest("GET", msixURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		rs.SetState(checklist.StepError)
		return "", false
	}
	defer resp.Body.Close()

	f, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

	writer := writerWrapper{
		Out:        f,
		TotalBytes: fileSize,
		Progress: func(current, total int64) {
			if total > 0 {
				var pp int8 = int8((float64(current) / float64(total)) * 100.0)
				rs.SetProgressPercentage(&pp)
			}
		},
	}
	io.Copy(&writer, resp.Body)

	if writer.Hash() != expectedSha512Sum {
		return "", false
	}

	rs.SetState(checklist.StepSuccess)
	return fileName, true
}

func installMSIXFromDownloadsFolder(msixPath string, doIOwnIt bool, rs checklist.RunStep) {
	runtime.LockOSThread()
	winrt.Initialize()
	// defer winrt.Uninitialize()
	rs.SetState(checklist.StepInProgress)

	msixURI := fmt.Sprint("file://", strings.ReplaceAll(msixPath, "\\", "/"))
	log.Println("MSIX", msixURI)
	uri := winrt.NewUri_CreateUri(msixURI)
	log.Println("MSIX URL:", msixURI, "abs", uri.Get_AbsoluteUri())

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

	var pp int8
	Ro_AwaitWithProgress(
		op,
		func(
			_ *winrt.IAsyncOperationWithProgress[*management.IDeploymentResult, management.DeploymentProgress],
			progress management.DeploymentProgress,
		) error {
			log.Println("Percentage:", progress.Percentage)
			pp = int8(progress.Percentage)
			rs.SetProgressPercentage(&pp)
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
						if doIOwnIt {
							os.Remove(msixPath)
						}
						log.Println("It worked")
					}
				}
			}
			return nil
		},
	)
	rs.SetState(checklist.StepSuccess)
}

func runInstalledApp(protocolHandler string, rs checklist.RunStep) {
	uri := protocolHandler + "://"

	rs.SetMessage("Running...")
	fmt.Println("Ran", uri, exec.Command("cmd", "/c", "start", uri).Run())
	rs.SetState(checklist.StepSuccess)
}

func isSilent() bool {
	for _, i := range os.Args {
		s := strings.ToLower(i)
		if s == "/q" || s == "/s" || s == "--silent" || s == "--quiet" {
			return true
		}
	}

	return false
}

func main() {
	installTarget := config.TargetProduct

	if installTarget == "" {
		installTarget = "Notion Stg"
		log.Println("No target set")
		// os.Exit(1)
	}

	var arch = channels.ArchAmd64
	if IsArm64() {
		arch = channels.ArchArm64
	}

	app := channels.DesktopProduct{
		ProductName:  installTarget,
		Architecture: arch,
	}

	if !IsRunningElevated() {
		RunElevated()
		return
	}

	cl := checklist.NewGioChecklist("Installing " + app.ProductName)

	msixstep := cl.AddStep("Finding latest app version online")
	stopstep := cl.AddStep("Stop current app")
	uninstallstep := cl.AddStep("Uninstall old app")
	installstep := cl.AddStep("Install app")
	runappstep := cl.AddStep("Run installed app")

	go func() {
		feed := channels.DesktopProductFeeds[app]
		msixUrl, fileSize, sha512, err := findLatestMSIXUpdate(feed.YamlFeed, msixstep)
		if err == nil {
			msixPath, doIOwnIt := downloadMSIXToDownloadsFolder(msixUrl, fileSize, sha512, msixstep)

			if msixPath != "" {
				stopAppIfRunning(app.ProductName, stopstep)
				uninstallWin32AppIfInstalled(app.ProductName, uninstallstep)
				installMSIXFromDownloadsFolder(msixPath, doIOwnIt, installstep)
				runInstalledApp(feed.Protocol, runappstep)
			} else {
				stopstep.SetState(checklist.StepSkipped)
				uninstallstep.SetState(checklist.StepSkipped)
				installstep.SetState(checklist.StepSkipped)
				runappstep.SetState(checklist.StepSkipped)
			}
		}

		cl.Finish()
	}()
	cl.Start()

	log.Println("Press key to get out")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadRune()
}
