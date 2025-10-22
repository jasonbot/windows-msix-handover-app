package main

import (
	"log"

	management "github.com/jasonbot/msix-installer-bootstrap/management"
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {
	winrt.Initialize()

	pm := management.NewPackageManager()
	for it := pm.FindPackages().First(); it.Get_HasCurrent() != false; it.MoveNext() {
		current := it.Get_Current()
		id := current.Get_Id()
		log.Println("Package", id.Get_Name())
	}
}
