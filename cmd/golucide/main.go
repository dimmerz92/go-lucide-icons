package main

import (
	"fmt"
	"os"

	"github.com/dimmerz92/go-templ-lucide/internal"
)

const INPUT = "./lucide/icons"
const HTML_OUTPUT = "./pkg/html/icons"
const TEMPL_OUTPUT = "./pkg/templ/icons"

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
	}

	switch args[1] {
	case "add":
		if len(args) < 4 {
			printUsage()
		}

		if internal.ReturnIcon(args[2:]) < 0 {
			printUsage()
		}

	case "sync":
		internal.SyncFiles(INPUT, HTML_OUTPUT, TEMPL_OUTPUT)

	case "test":
		internal.TestServer(args[2:])

	case "help", "-h", "--help":
		fmt.Print(USAGE)

	default:
		printUsage()
	}
}

const USAGE = `
	go-lucide - a port of lucide icons for Go developers

	USAGE: golucide <command> [<args>...]

	COMMANDS:

	add [-o ouput_file] <templ | html> <icon name (kebab-case)>
	*
	* Adds a templ or html icon template to your project.
	* The icon template will be generated in the directory the command was run
	* from if an output file is not specified.

	sync
	*
	* Syncs new icons from the ./lucide/icons directory to the relevant html or
	* templ directory in the ./pkg directory.

	test [-p <port>]
	*
	* Runs the test server to view and test the rendering of icons.

`

func printUsage() {
	fmt.Print(USAGE)
	os.Exit(1)
}
