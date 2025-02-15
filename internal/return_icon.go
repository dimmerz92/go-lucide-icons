package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dimmerz92/go-templ-lucide/icons"
)

func ReturnIcon(args []string) int {
	if len(args) < 2 {
		return -1
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	switch args[0] {
	case "html":
		file, err := icons.GetHtmlFile(args[1] + ".html")
		if err != nil {
			log.Fatalf("file not found: %v", err)
		}

		err = os.WriteFile(filepath.Join(cwd, args[1]+".html"), file, 0644)
		if err != nil {
			log.Fatalf("failed to write file to output: %v", err)
		}

	case "templ":
		file, err := icons.GetTemplFile(args[1] + ".templ")
		if err != nil {
			log.Fatalf("file not found: %v", err)
		}

		err = os.WriteFile(filepath.Join(cwd, args[1]+".templ"), file, 0644)
		if err != nil {
			log.Fatalf("failed to write file to output: %v", err)
		}

	default:
		return -1
	}

	return 0
}
