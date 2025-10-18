package main

import (
	"flag"
	"os"
	"path/filepath"

	htmlicons "github.com/dimmerz92/go-lucide-icons/html-icons"
	templicons "github.com/dimmerz92/go-lucide-icons/templ-icons"
	"github.com/fatih/color"
)

const PERMS = 0600

func main() {
	args := os.Args
	if len(args) < 2 {
		println(help)
		os.Exit(1)
	}

	switch args[1] {
	case "html":
		if len(args) < 3 {
			color.Red("error: icon name required\n")
			println(help)
			os.Exit(1)
		}

		f := flag.NewFlagSet("templ", flag.ContinueOnError)
		output := f.String("out", ".", "the output directory")
		f.Parse(args[3:])

		file, err := htmlicons.GetHtmlFile(args[2])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		err = os.WriteFile(filepath.Join(*output, args[2]+".html"), file, PERMS)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		color.Green("%s.html save to %s", args[2], *output)

	case "templ":
		if len(args) < 3 {
			color.Red("error: icon name required\n")
			println(help)
			os.Exit(1)
		}

		f := flag.NewFlagSet("templ", flag.ContinueOnError)
		output := f.String("out", ".", "the output directory")
		f.Parse(args[3:])

		file, err := templicons.GetTemplFile(args[2])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		err = os.WriteFile(filepath.Join(*output, args[2]+".templ"), file, PERMS)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		color.Green("%s.templ save to %s", args[2], *output)

	case "help":
		fallthrough

	default:
		if args[1] != "help" {
			color.Red("%s: not a valid command\n", args[1])
		}
		println(help)
	}
}

var help = color.YellowString("USAGE:\n") +
	color.WhiteString("\tgolucide ") + color.BlueString("<COMMAND> ") + color.MagentaString("[OPTIONS]\n\n") +
	color.YellowString("COMMANDS:\n") +
	color.BlueString("\thtml ") + color.CyanString("<icon name> ") + color.MagentaString("[options]\n") +
	color.WhiteString("\tGenerates a html lucide icon template.\n") +
	color.MagentaString("\t-out") + color.WhiteString(" directory to save the generated icon. Defaults to .\n\n") +
	color.BlueString("\ttempl ") + color.CyanString("<icon name> ") + color.MagentaString("[options]\n") +
	color.WhiteString("\tGenerates a templ lucide icon template.\n") +
	color.MagentaString("\t-out") + color.WhiteString(" directory to save the generated icon. Defaults to .\n\n") +
	color.BlueString("\thelp\n") +
	color.WhiteString("\tPrints help text for golucide.\n")
