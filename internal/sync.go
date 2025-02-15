package internal

import (
	"log"
	"os/exec"
	"runtime"
	"sync"
)

// SyncFiles runs through the list of source files and generates templ and html
// icons and puts them in the target if not already there.
func SyncFiles(src, target string) {
	newFiles, err := getNewFiles(src, target)
	if err != nil {
		log.Fatalln(err)
	}

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
		go processFile(target, &wg, tasks, errors)
	}

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

// processFile encapsulates the html and templ icon generation methods.
func processFile(target string, wg *sync.WaitGroup, tasks <-chan string, errors chan<- error) {
	defer wg.Done()

	for file := range tasks {
		if err := generateHtmlIcon(file, target); err != nil {
			errors <- err
		}

		if err := generateTemplIcon(file, target); err != nil {
			errors <- err
		}
	}
}

// templGenerate runs the templ formatter on the generated templ files and then
// runs the templ generator on the generated templ files to generate Go code.
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
