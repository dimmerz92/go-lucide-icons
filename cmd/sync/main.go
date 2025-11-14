package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dimmerz92/go-lucide-icons/internal"
)

const (
	LUCIDE_ICONS_DIR = "./lucide/icons"
	HTML_ICONS_DIR   = "./htmlicons/templates"
	TEMPL_ICONS_DIR  = "./templicons/icons"
)

func main() {
	svgIcons := internal.FileSet(LUCIDE_ICONS_DIR, ".svg")
	htmlWant := internal.DiffFileSet(svgIcons, internal.FileSet(HTML_ICONS_DIR, ".html"))
	templWant := internal.DiffFileSet(svgIcons, internal.FileSet(TEMPL_ICONS_DIR, ".templ"))

	for svg := range svgIcons {
		_, inHtml := htmlWant[svg]
		_, inTempl := templWant[svg]

		var file []byte
		var err error

		if inHtml || inTempl {
			file, err = os.ReadFile(filepath.Join(LUCIDE_ICONS_DIR, svg+".svg"))
			if err != nil {
				panic(err)
			}
		}

		if inHtml {
			err = internal.ToHTML(svg, string(file), HTML_ICONS_DIR)
			if err != nil {
				panic(err)
			}
		}

		if inTempl {
			err = internal.ToTempl(svg, string(file), TEMPL_ICONS_DIR)
			if err != nil {
				panic(err)
			}
		}
	}

	_, err := exec.Command("templ", "fmt", TEMPL_ICONS_DIR).Output()
	if err != nil {
		panic(err)
	}

	_, err = exec.Command("templ", "generate").Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Finished: %d html files & %d templ files", len(htmlWant), len(templWant))
}
