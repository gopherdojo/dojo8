package main

import (
	"dojo8/kadai1/wataboru/imageconverter"

	"flag"
	"fmt"
	"os"
)

const (
	// ExitCodeSuccess is the exit code on success
	ExitCodeSuccess int = iota
	// ExitCodeError is the exit code when failed
	ExitCodeError
	// ExitCodeError is the exit code when failed
	ExitCodeInvalidDirectoryError
)

var (
	args imageconverter.Args
)

func init() {
	flag.StringVar(&args.Directory, "dir", "", "Input target Directory.\n  ex) `--dir=./convert_image`")
	flag.StringVar(&args.Directory, "d", "", "Input target Directory. (short)")
	flag.StringVar(&args.BeforeExtension, "before", "jpg", "Input extension before conversion.\n  ex) `--before=png`")
	flag.StringVar(&args.BeforeExtension, "b", "jpg", "Input extension before conversion. (short)")
	flag.StringVar(&args.AfterExtension, "after", "png", "Input extension after conversion.\n  ex) `--after=jpg`")
	flag.StringVar(&args.AfterExtension, "a", "png", "Input extension after conversion. (short)")
	flag.Parse()
}

func run() int {
	if args.Directory == "" {
		fmt.Fprintln(os.Stderr, "Input target Directory.\n  ex) `--dir=./convert_image`")
		return ExitCodeInvalidDirectoryError
	}

	if _, err := os.Stat(args.Directory); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Target directory is not found.")
		return ExitCodeInvalidDirectoryError
	}

	if err := imageconverter.Convert(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitCodeError
	}

	return ExitCodeSuccess
}

func main() {
	os.Exit(run())
}
