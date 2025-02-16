package gltempl

import (
	"embed"
	"io/fs"
)

//go:embed icons/*.templ
var templEmbed embed.FS

func GetTemplFile(name string) ([]byte, error) {
	return fs.ReadFile(templEmbed, "icons/"+name)
}
