package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo8/kadai1/tanaka0325/imgconv"
)

func main() {
	if err := imgconv.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
