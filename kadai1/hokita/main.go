package main

import (
	"flag"
	"fmt"
	"imgconv/imgconv"
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
	os.Exit(run())
}

func run() int {
	flag.Parse()
	dir := flag.Arg(0)

	err := imgconv.Call(dir, from, to)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	fmt.Println("Conversion finished!")
	return ExitCodeOk
}
