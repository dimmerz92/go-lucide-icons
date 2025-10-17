package htmlicons

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed templates/*.html
var templates embed.FS

func GetHtmlFile(name string) ([]byte, error) {
	return fs.ReadFile(templates, path.Join("templates", name+".html"))
}
