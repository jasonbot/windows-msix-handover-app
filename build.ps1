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

foreach ($product in $Products) {
    ~/go/bin/go-winres simply --icon icon.ico --manifest gui --out cmd/pullpackage/rsrc --file-description="Installs the latest $($product) Desktop" --product-name="$($product) Installer" --copyright="Notion Labs, Inc." --file-version=$(Get-Date -Format yyyy.MM.dd.HHmm)
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
