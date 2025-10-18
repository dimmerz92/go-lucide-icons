package htmlicons

import (
	"embed"
	"html/template"
	"io/fs"
	"path"
)

//go:embed templates/*.html
var templates embed.FS

func GetHtmlFile(name string) ([]byte, error) {
	return fs.ReadFile(templates, path.Join("templates", name+".html"))
}

// AddLucideIcons combines the lucide icons with the given template.
func AddLucideIcons(tpls *template.Template) error {
	_, err := tpls.ParseFS(templates, "templates/*.html")
	return err
}
