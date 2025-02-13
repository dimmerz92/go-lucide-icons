package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const ICON_DIR = "./icons/"
const LUCIDE_DIR = "./lucide/icons/"
const GOFUNC = `package icons

templ %s(props IconProps) {
    %s
		if props.ID != "" {
			id={ props.ID }
		}
		if props.Class != "" {
			class={ props.Class }
		}
		if props.Style != "" {
			style={ props.Style }
		}
		{ props.Attributes... }
    >
    %s
}`

var existingFiles = map[string]bool{}
var newFiles = []string{}

func main() {
	// get a list of icons that we already have and store in memory
	iconDir, err := os.ReadDir(ICON_DIR)
	if err != nil {
		log.Fatalf("failed to read icons directory: %v", err)
	}

	for _, file := range iconDir {
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		existingFiles[name] = true
	}

	// get a list of svg icons from the cloned lucide icons repo
	lucideDir, err := os.ReadDir(LUCIDE_DIR)
	if err != nil {
		log.Fatalf("failed to read lucide icons directory: %v", err)
	}

	for _, file := range lucideDir {
		// skip if not an svg
		if !strings.HasSuffix(file.Name(), ".svg") {
			continue
		}

		// skip if already converted
		name := strings.TrimSuffix(file.Name(), ".svg")
		if _, ok := existingFiles[name]; ok {
			continue
		}

		newFiles = append(newFiles, name)
	}

	// process images
	var wg sync.WaitGroup
	nProc := runtime.NumCPU()
	errors := make(chan error, nProc)
	tasks := make(chan string, len(newFiles))

	for _, file := range newFiles {
		tasks <- file
	}
	close(tasks)

	for i := 0; i < nProc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for name := range tasks {
				// kebab case to pascal case
				funcName := kebabToPascalCase(name)

				// convert svg to templ template string
				err := processImage(name, funcName)
				if err != nil {
					errors <- err
					continue
				}

			}
		}()
	}

	// handle any errors
	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		log.Printf("failed to convert icon: %v", err)
	}

	// run templ fmt and generate
	if err := templGenerate(); err != nil {
		log.Fatalf("failed to run templ commands: %v", err)
	}
}

// converts a string from kebab case to pascal case
func kebabToPascalCase(s string) string {
	var buf strings.Builder
	for _, sub := range strings.Split(s, "-") {
		buf.WriteString(strings.ToUpper(string(sub[:1])) + sub[1:])
	}
	return buf.String()
}

// opens the given filename and converts the svg to a templ component
func processImage(name, funcname string) error {
	// read the file
	content, err := os.ReadFile(LUCIDE_DIR + name + ".svg")
	if err != nil {
		return err
	}

	// split on the opening svg tag
	parts := strings.SplitN(string(content), ">", 2)
	if len(parts) != 2 {
		return fmt.Errorf("Failed to convert svg: %s", name)
	}

	// interpolate parts into templ template
	templ := fmt.Sprintf(GOFUNC, funcname, parts[0], parts[1])

	// save templ file to icons folder
	err = os.WriteFile(ICON_DIR+name+".templ", []byte(templ), 0644)
	if err != nil {
		return err
	}

	return nil
}

func templGenerate() error {
	cmd := exec.Command("templ", "fmt", "icons")
	if _, err := cmd.Output(); err != nil {
		return err
	}

	cmd = exec.Command("templ", "generate")
	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
