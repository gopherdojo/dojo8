package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOk    = 0
	ExitCodeError = 1
)

var from, to string

func init() {
	flag.StringVar(&from, "from", "jpg", "Conversion source extension.")
	flag.StringVar(&to, "to", "png", "Conversion target extension.")
}

func main() {
	flag.Parse()
	exitCode := run(flag.Arg(0))
	os.Exit(exitCode)
}

func run(arg string) int {
	err := convert(arg, from, to)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	fmt.Println("Conversion finished!")
	return ExitCodeOk
}
