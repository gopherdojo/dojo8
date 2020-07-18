package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo8/kadai2/kynefuk/cli"
)

func main() {
	var directory string
	var from string
	var to string
	flag.StringVar(&directory, "directory", "./", "directory")
	flag.StringVar(&directory, "d", "./", "directory(short)")
	flag.StringVar(&from, "from", "jpg", "from format")
	flag.StringVar(&from, "f", "jpg", "from format(short)")
	flag.StringVar(&to, "to", "png", "to format")
	flag.StringVar(&to, "t", "png", "to format(short)")
	flag.Parse()

	command := cli.NewCommand(os.Stdout, os.Stderr)

	args := cli.NewArgs(directory, from, to)
	if err := args.Validate(); err != nil {
		fmt.Fprintf(command.ErrStream, "error: %s\n", err)
		os.Exit(cli.ExitCodeError)
	}

	os.Exit(command.Run(args))
}
