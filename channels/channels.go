package channels

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

type DesktopFeed struct {
	YamlFeed string
	AppID    string
	Protocol string
}

var DesktopProductFeeds map[DesktopProduct]DesktopFeed

func init() {
	DesktopProductFeeds = map[DesktopProduct]DesktopFeed{
		DesktopProduct{
			ProductName:  "Notion",
			Architecture: ArchArm64,
		}: {
			YamlFeed: "https://desktop-release.notion-static.com/arm64-msix.yml",
			AppID:    "com.notion.app.desktop.notion",
			Protocol: "notion",
		},
		DesktopProduct{
			ProductName:  "Notion Dev",
			Architecture: ArchArm64,
		}: {
			YamlFeed: "https://dev-desktop-release.s3.us-west-2.amazonaws.com/arm64-msix.yml",
			AppID:    "com.notion.app.desktop.notiondev",
			Protocol: "notiondev",
		},
		DesktopProduct{
			ProductName:  "Notion Stg",
			Architecture: ArchArm64,
		}: {
			YamlFeed: "https://stg-desktop-release.s3.us-west-2.amazonaws.com/arm64-msix.yml",
			AppID:    "com.notion.app.desktop.notionstg",
			Protocol: "notionstg",
		},
		DesktopProduct{
			ProductName:  "Notion",
			Architecture: ArchAmd64,
		}: {
			YamlFeed: "https://desktop-release.notion-static.com/msix.yml",
			AppID:    "com.notion.app.desktop.notion",
			Protocol: "notion",
		},
		DesktopProduct{
			ProductName:  "Notion Dev",
			Architecture: ArchAmd64,
		}: {
			YamlFeed: "https://dev-desktop-release.s3.us-west-2.amazonaws.com/msix.yml",
			AppID:    "com.notion.app.desktop.notiondev",
			Protocol: "notiondev",
		},
		DesktopProduct{
			ProductName:  "Notion Stg",
			Architecture: ArchAmd64,
		}: {
			YamlFeed: "https://stg-desktop-release.s3.us-west-2.amazonaws.com/msix.yml",
			AppID:    "com.notion.app.desktop.notionstg",
			Protocol: "notionstg",
		},
	}
}
