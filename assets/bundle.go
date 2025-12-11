package assets

import (
	"embed"

	"github.com/Kwynto/vql-prototype/internal/manage/webmanage"
)

var (
	//go:embed ui/html
	emUiHtmlDir embed.FS

	//go:embed ui/static
	emUiStaticDir embed.FS
)

func init() {
	// Set UI
	webmanage.SetUiDirs(&emUiHtmlDir, &emUiStaticDir)
}
