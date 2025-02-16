package internal

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/dimmerz92/go-templ-lucide/pkg/html/icons"
)

func TestServer(args []string) int {
	cmd := flag.NewFlagSet("test", flag.ExitOnError)
	port := cmd.Int("p", 8000, "use a port between 1024 and 65535")
	cmd.Parse(args)

	if *port < 1024 || *port > 65535 {
		fmt.Printf(
			"port: %d out of range, use a port between 1024 and 65535",
			*port,
		)
		os.Exit(1)
	}

	switch args[len(args)-1] {
	case "html":
		data := struct {
			Medium []template.HTMLAttr
			Large  []template.HTMLAttr
			XLarge []template.HTMLAttr
		}{
			Medium: []template.HTMLAttr{`style="height: 2rem; width: 2rem"`},
			Large:  []template.HTMLAttr{`style="height: 3rem; width: 3rem"`},
			XLarge: []template.HTMLAttr{`style="height: 4rem; width: 4rem"`},
		}
		templates := template.Must(template.ParseFiles("internal/index.html"))
		icons.AddLucideIcons(templates)
		http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			templates.Execute(w, data)
		})

	case "templ":
		http.Handle("GET /", templ.Handler(index()))

	default:
		return -1
	}

	fmt.Printf("Listening on :%d", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)

	return 0
}
