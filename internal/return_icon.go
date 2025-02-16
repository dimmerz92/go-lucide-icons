package internal

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	glhtml "github.com/dimmerz92/go-templ-lucide/pkg/html"
	gltempl "github.com/dimmerz92/go-templ-lucide/pkg/templ"
)

func ReturnIcon(args []string) int {
	if len(args) < 2 {
		return -1
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	cmd := flag.NewFlagSet("add", flag.ExitOnError)
	output := cmd.String("o", cwd, "use a port between 1024 and 65535")
	cmd.Parse(args)

	icon := len(args) - 1

	switch args[len(args)-2] {
	case "html":
		file, err := glhtml.GetHtmlFile(args[icon] + ".html")
		if err != nil {
			log.Fatalf("file not found: %v", err)
		}

		err = os.WriteFile(filepath.Join(*output, args[icon]+".html"), file, 0644)
		if err != nil {
			log.Fatalf("failed to write file to output: %v", err)
		}

	case "templ":
		file, err := gltempl.GetTemplFile(args[icon] + ".templ")
		if err != nil {
			log.Fatalf("file not found: %v", err)
		}

		err = os.WriteFile(filepath.Join(*output, args[icon]+".templ"), file, 0644)
		if err != nil {
			log.Fatalf("failed to write file to output: %v", err)
		}

	default:
		return -1
	}

	return 0
}
