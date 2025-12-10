$Products = @(
    "Notion",
    "Notion Dev",
    "Notion Stg"
)

$Architectures = @(
    "amd64",
    "arm64"
)

go install github.com/tc-hib/go-winres@latest
~/go/bin/go-winres simply --icon icon.ico --manifest gui --out cmd/pullpackage/rsrc

foreach ($product in $Products) {
    foreach ($arch in $Architectures) {
        $env:GOARCH = $arch
        $exeFile = "Install $($product) as MSIX ($arch).exe"
        write-host "Making", $exeFile
        go build -ldflags "-X 'github.com/jasonbot/windows-msix-handover-app/config.TargetProduct=$($product)'" -o "$($exeFile)" ./cmd/pullpackage
        if ($arch -eq "amd64") {
            upx $exeFile
        }
    }
}
