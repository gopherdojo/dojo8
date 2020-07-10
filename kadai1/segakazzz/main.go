package main

import (
	"os"

	"github.com/gopherdojo/dojo8/kadai1/segakazzz/imgconv"
)

func main() {
	err := imgconv.RunConverter()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
