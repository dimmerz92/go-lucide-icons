package icons

import (
	"embed"
	"io/fs"
)

//go:embed *.html
var htmlEmbed embed.FS

//go:embed *.templ
var templEmbed embed.FS

func GetHtmlFile(name string) ([]byte, error) {
	return fs.ReadFile(htmlEmbed, name)
}

func GetTemplFile(name string) ([]byte, error) {
	return fs.ReadFile(templEmbed, name)
}
