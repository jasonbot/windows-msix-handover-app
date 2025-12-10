# windows-msix-handover-app
A platypus system for moving from .exe installers to .msix


```shell
go get github.com/akavel/rsrc
rsrc -manifest pullpackage.exe.manifest -o app.syso
go build -o pullpackage.exe ./cmd/pullpackage
```
