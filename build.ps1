$Products = @(
    "Notion",
    "Notion Dev",
    "Notion Stg"
)

go install github.com/tc-hib/go-winres@latest

$arch = "amd64"
$env:GOARCH = $arch

foreach ($product in $Products) {
    ~/go/bin/go-winres simply --icon icon.png --manifest gui --out cmd/pullpackage/rsrc --file-description="Installs the latest $($product) Desktop" --product-name="$($product) Installer" --copyright="Notion Labs, Inc." --file-version=git-tag --product-version=$(Get-Date -Format yyyy.MM.dd.HHmm)
    $exeFile = "Install Latest $($product) Desktop.exe"
    write-host "Making", $exeFile
    go build -ldflags "-X 'github.com/jasonbot/windows-msix-handover-app/config.TargetProduct=$($product)'" -o "$($exeFile)" ./cmd/pullpackage
    upx $exeFile
}
