package glhtml

import (
	"embed"
	"io/fs"
)

//go:embed icons/*.html
var htmlEmbed embed.FS

func GetHtmlFile(name string) ([]byte, error) {
	return fs.ReadFile(htmlEmbed, "icons/"+name)
}
