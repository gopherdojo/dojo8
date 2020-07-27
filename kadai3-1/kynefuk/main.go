package main

import (
	"flag"
	"os"

	"github.com/gopherdojo/dojo8/kadai3-1/kynefuk/cmd"
)

func main() {
	var limit int
	flag.IntVar(&limit, "limit", 60, "Enter a time limit")
	flag.IntVar(&limit, "l", 60, "Enter a time limit")
	flag.Parse()

	cli := cmd.NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(limit))
}
