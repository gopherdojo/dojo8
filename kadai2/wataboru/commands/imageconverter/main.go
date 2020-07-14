package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/gopherdojo/dojo8/kadai1/wataboru/imageconverter"
)

const (
	// ExitCodeSuccess is the exit code on success
	ExitCodeSuccess = 0
	// ExitCodeError is the exit code when failed
	ExitCodeError = 1
	// ExitCodeError is the exit code when failed
	ExitCodeInvalidDirectoryError = 2
)

var (
	args         imageconverter.Args
	osStat       func(name string) (os.FileInfo, error)
	osIsNotExist func(err error) bool
	imgconv      func(args imageconverter.Args) error
)

func init() {
	testing.Init()
	flag.StringVar(&args.Directory, "dir", "", "Input target Directory.\n  ex) `--dir=./convert_image`")
	flag.StringVar(&args.Directory, "d", "", "Input target Directory. (short)")
	flag.StringVar(&args.BeforeExtension, "before", "jpg", "Input extension before conversion.\n  ex) `--before=png`")
	flag.StringVar(&args.BeforeExtension, "b", "jpg", "Input extension before conversion. (short)")
	flag.StringVar(&args.AfterExtension, "after", "png", "Input extension after conversion.\n  ex) `--after=jpg`")
	flag.StringVar(&args.AfterExtension, "a", "png", "Input extension after conversion. (short)")
	flag.Parse()

	osStat = os.Stat
	osIsNotExist = os.IsNotExist
	imgconv = imageconverter.Convert

}

func run() int {
	if args.Directory == "" {
		fmt.Fprintln(os.Stderr, "Input target Directory.\n  ex) `--dir=./convert_image`")
		return ExitCodeInvalidDirectoryError
	}

	if _, err := osStat(args.Directory); osIsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Target directory is not found.")
		return ExitCodeInvalidDirectoryError
	}

	if err := imgconv(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitCodeError
	}

	return ExitCodeSuccess
}

func main() {
	os.Exit(run())
}
