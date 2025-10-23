package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	management "github.com/jasonbot/windows-msix-handover-app/management"
	"github.com/shirou/gopsutil/process"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
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

func downloadMSIXToDownloadsFolder(msixURL string) string {
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

	num := 1

	_, err := os.Stat(fileName)
	for !errors.Is(err, os.ErrNotExist) {
		num += 1
		fileName = fmt.Sprintf("%v (%v)%v", fileBase, num, ext)
		_, err = os.Stat(fileName)
	}

	req, _ := http.NewRequest("GET", msixURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		return ""
	}
	defer resp.Body.Close()

	f, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

	/*
	   bar := progressbar.DefaultBytes(
	      resp.ContentLength,
	      "downloading",
	   )
	   io.MultiWriter(f, bar)
	*/

	io.Copy(f, resp.Body)

	return fileName
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
	op.Put_Progress(
		func(
			asyncInfo *winrt.IAsyncOperationWithProgress[
				*management.IDeploymentResult,
				management.DeploymentProgress,
			], progressInfo management.DeploymentProgress,
		) com.Error {
			_ = progressInfo.Percentage
			return com.OK
		})
	op.Put_Completed(
		func(
			asyncInfo *winrt.IAsyncOperationWithProgress[
				*management.IDeploymentResult,
				management.DeploymentProgress,
			],
			asyncStatus winrt.AsyncStatus,
		) com.Error {
			if asyncStatus == winrt.AsyncStatus_Completed || asyncStatus == winrt.AsyncStatus_Error {
				if r := asyncInfo.GetResults(); r != nil {
					_ = r.Get_ErrorText()
				}
			}
			win32.PostThreadMessage(com.GetContext().TID, win32.WM_QUIT, 0, 0)
			return com.OK
		})
	com.MessageLoop()
}

func main() {
	appName := "Notion Dev"

	stopAppIfRunning(appName)
	uninstallAppIfInstalled(appName)
}
