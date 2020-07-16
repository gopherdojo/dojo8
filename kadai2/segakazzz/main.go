package main

import (
	"flag"
	"os"

	"github.com/gopherdojo/dojo8/kadai2/segakazzz/imgconv"
)

func main() {
	var (
		dir = flag.String("d", ".", "Indicate directory to convert")
		in  = flag.String("i", "jpg", "Indicate input image file's extension")
		out = flag.String("o", "png", "Indicate output image file's extension")
	)

	flag.Parse()
	err := imgconv.RunConverter(dir, in, out)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
