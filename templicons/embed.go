package templicons

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed icons/*.templ
var templates embed.FS

func GetTemplFile(name string) ([]byte, error) {
	return fs.ReadFile(templates, path.Join("icons", name+".templ"))
}
