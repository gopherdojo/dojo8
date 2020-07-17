package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo8/kadai1/skmatz/imgconv"
)

func run() error {
	var opt imgconv.Options

	opt.Dir = flag.String("dir", "", "Path to the target directory")
	opt.From = flag.String("from", "jpg", "Image extension before conversion")
	opt.To = flag.String("to", "png", "Image extension after conversion")
	opt.Verbose = flag.Bool("verbose", false, "Show conversion logs")

	flag.Parse()

	cvt, err := imgconv.NewImageConverter(*opt.From, *opt.To)
	if err != nil {
		return err
	}

	if err := cvt.ConvertAll(*opt.Dir, *opt.Verbose); err != nil {
		return err
	}

	return nil
}

func main() {
	exitCode := 0

	if err := run(); err != nil {
		exitCode = 1
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitCode)
}
