package internal

import (
	"log"
	"os/exec"
	"runtime"
	"sync"
)

// SyncFiles runs through the list of source files and generates templ and html
// icons and puts them in the target if not already there.
func SyncFiles(src, htmlTarget, templTarget string) {
	newFiles, err := getNewFiles(src, htmlTarget, templTarget)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	nProc := runtime.NumCPU()
	errors := make(chan error, nProc)
	htmlTasks := make(chan string, len(newFiles["html"]))
	templTasks := make(chan string, len(newFiles["templ"]))

	for _, file := range newFiles["html"] {
		htmlTasks <- file
	}
	close(htmlTasks)

	for _, file := range newFiles["templ"] {
		templTasks <- file
	}
	close(templTasks)

	for i := 0; i < nProc; i++ {
		wg.Add(1)
		go processFile(
			htmlTarget,
			templTarget,
			&wg,
			htmlTasks,
			templTasks,
			errors,
		)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		log.Printf("failed to convert icon: %v", err)
	}

	// run templ fmt and generate
	if err := templGenerate(templTarget); err != nil {
		log.Fatalf("failed to run templ commands: %v", err)
	}
}

// processFile encapsulates the html and templ icon generation methods.
func processFile(
	htmlTarget,
	templTarget string,
	wg *sync.WaitGroup,
	htmlTasks, templTasks <-chan string,
	errors chan<- error,
) {
	defer wg.Done()

	for file := range htmlTasks {
		if err := generateHtmlIcon(file, htmlTarget); err != nil {
			if err != nil {
				errors <- err
			}
		}

		for file := range templTasks {
			if err := generateTemplIcon(file, templTarget); err != nil {
				errors <- err
			}
		}

	}
}

// templGenerate runs the templ formatter on the generated templ files and then
// runs the templ generator on the generated templ files to generate Go code.
func templGenerate(target string) error {
	cmd := exec.Command("templ", "fmt", target)
	if _, err := cmd.Output(); err != nil {
		return err
	}

	cmd = exec.Command("templ", "generate")
	if _, err := cmd.Output(); err != nil {
		return err
	}

	return nil
}
