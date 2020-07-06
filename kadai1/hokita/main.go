package main

import (
	"fmt"
	"imgconv/imgconv"
	"os"
)

const (
	ExitCodeOk int = iota
	ExitCodeError
)

func main() {
	os.Exit(run())
}

func run() int {
	err := imgconv.Call()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	fmt.Println("Conversion finished!")
	return ExitCodeOk
}
