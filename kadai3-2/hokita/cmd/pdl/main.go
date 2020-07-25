package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo8/kadai3-2/hokita/pdl"
)

const (
	ExitCode    = 0
	ExitCodeErr = 1
)

var proc int
var dir string

func init() {
	flag.IntVar(&proc, "proc", 10, "split ratio to download file")
	flag.StringVar(&dir, "dir", "testdata", "output file")
}

func main() {
	flag.Parse()
	exitCode := run(flag.Arg(0))
	os.Exit(exitCode)
}

func run(url string) int {
	cli, err := pdl.New(proc, url, dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v\n", err)
		return ExitCodeErr
	}

	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v\n", err)
		return ExitCodeErr
	}

	return ExitCode
}
