package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo8/kadai1/tanaka0325/imgconv"
)

var options imgconv.Options
var args imgconv.Args

func init() {
	options.From = flag.String("f", "jpg", "file extension before convert")
	options.To = flag.String("t", "png", "file extension after convert")
	options.DryRun = flag.Bool("n", false, "dry run")
	flag.Parse()

	args = flag.Args()
}

func main() {
	if err := imgconv.Run(options, args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
