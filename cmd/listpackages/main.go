package main

import (
	"log"
	"runtime"

	management "github.com/jasonbot/msix-installer-bootstrap/management"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {
	runtime.LockOSThread()
	winrt.Initialize()
	defer winrt.Uninitialize()

	pm := management.NewPackageManager()
	if pm.IUnknown.GetIUnknown() == nil {
		log.Println("Ooof")
		return
	}
	fp, err := pm.FindPackages()
	if err != nil {
		log.Println("Ooof", err)
		return
	}
	for it := fp.First(); it != nil && it.Get_HasCurrent() != false; it.MoveNext() {
		current := it.Get_Current()
		id := current.Get_Id()
		log.Println("Package", id.Get_Name())
	}

	com.MessageLoop()
}
