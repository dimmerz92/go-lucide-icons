package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const GOFUNC = `package icons

templ %s(class ...string) {
    %s
		class={ class }
    >
    %s
}
`

func main() {
	f, _ := os.Create("test.txt")
	f.Write([]("hello world"))
	f.Close()
	fmt.Println("test succeeded")
}

// Encapsulates the svg files as a templ component and adds variadic class html
// attributes. Reads from the imgs folder and generates components to the icons
// folder. Calls templ fmt on the new icons for proper formatting.
func main1() {
	// read the imgs folder
	files, err := os.ReadDir("./imgs")
	if err != nil {
		log.Fatalf("Failed to read img directory: %s", err)
	}

	// get the names of all svg images that are not already converted
	imgs := []string{}
	for _, file := range files {
		var img string
		if strings.HasSuffix(file.Name(), ".svg") {
			img = strings.TrimSuffix(file.Name(), ".svg")
		}

		if img == "" {
			continue
		}

		_, err := os.Stat("./icons/" + img + ".templ")
		if !errors.Is(err, os.ErrNotExist) {
			continue
		}

		imgs = append(imgs, img)
	}

	// convert all svg images to templ component
	for _, img := range imgs {
		var funcname string

		// create a pascal case func name
		for _, title := range strings.Split(img, "-") {
			funcname += cases.Title(language.English).String(title)
		}

		// generate the templ component text
		templ := convertToTempl("./imgs/"+img+".svg", funcname)

		// save the output to a templ file
		err := os.WriteFile("./icons/"+img+".templ", []byte(templ), 0644)
		if err != nil {
			log.Printf("Failed to create templ component: %v", err)
			continue
		}
		log.Printf("Converted: %s", img)
	}

	// run templ fmt on the icons folder
	cmd := exec.Command("templ", "fmt", "icons")
	if _, err := cmd.Output(); err != nil {
		log.Printf("Failed to run templ fmt on the new components: %s", err)
	} else {
		log.Println("Successfully finished")
	}
}

// opens the given filename and converts the svg to a templ component
func convertToTempl(filename, funcname string) string {
	// read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
	}

	// split on the opening svg tag
	imgFragments := strings.SplitN(string(content), ">", 2)
	if len(imgFragments) != 2 {
		log.Printf("Failed to convert svg: %s", filename)
	}

	return fmt.Sprintf(GOFUNC, funcname, imgFragments[0], imgFragments[1])
}
