package main

import (
	"fmt"
	"os"

	"github.com/dimmerz92/go-templ-lucide/internal"
)

const INPUT = "./lucide/icons"
const OUTPUT = "./icons"

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
		internal.SyncFiles(INPUT, OUTPUT)

	case "test":
		internal.TestServer(args[2:])

	case "help", "-h", "--help":
		fmt.Printf(USAGE, INPUT, OUTPUT)

	default:
		printUsage()
	}
}

const USAGE = `
	go-lucide - a port of lucide icons for Go developers

	USAGE: golucide <command> [<args>...]

	COMMANDS:

	add <templ | html> <icon name (kebab-case)>
	*
	* Adds a templ or html icon template to your project.
	* Specify templ or html after the add command.
	* The icon template will be generated in the directory the command was run
	* from.

	sync [-i <input directory>] [-o <output directory>]
	*
	* Syncs new icons from the input to the output. If -i and/or -o flags are
	* not used, these values default to:
	*	-i %s
	*	-o %s

	test [-p <port>]
	*
	* Runs the test server to view and test the rendering of icons.

`

func printUsage() {
	fmt.Printf(USAGE, INPUT, OUTPUT)
	os.Exit(1)
}
