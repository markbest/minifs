package main

import (
	"flag"
	"fmt"
	"io"
	"github.com/markbest/minifs/command"
	"os"
	"strings"
	"text/template"
)

var commands = command.Commands

var usageTemplate = `
MiniFs: store billions of files and serve them fast!

Usage:
	minifs command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "minifs help [command]" for more information about a command.
`

var helpTemplate = `Usage: minifs {{.UsageLine}}
  {{.Long}}
`

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		for _, cmd := range commands {
			if len(args) >= 2 && cmd.Name() == args[1] && cmd.Run != nil {
				cmd.Flag.PrintDefaults()
			}
		}
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			if !cmd.Run(cmd, args) {
				fmt.Fprintf(os.Stderr, "\n")
				cmd.Flag.Usage()
				cmd.Flag.PrintDefaults()
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "minifs: unknown subcommand %q\nRun 'minifs help' for usage.\n", args[0])
}

// template executes the given template text on data, writing the result to w.
func tmp(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

// show default command usage
func usage() {
	fmt.Fprintf(os.Stderr, "use \"minifs [command]\"\n")
	os.Exit(2)
}

// help implements the 'help' command.
func help(args []string) {
	if len(args) == 0 {
		tmp(os.Stdout, usageTemplate, commands)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: minifs help command\n\nToo many arguments given\n")
		os.Exit(2)
	}

	arg := args[0]
	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmp(os.Stdout, helpTemplate, cmd)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q. Run 'minifs help'.\n", arg)
	os.Exit(2)
}