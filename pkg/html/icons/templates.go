package icons

import (
	"embed"
	"html/template"
)

//go:embed *.html
var iconFiles embed.FS

// AddLucideIcons combines the lucide icons with the given template.
func AddLucideIcons(templates *template.Template) {
	templates.ParseFS(iconFiles, "*.html")
}
