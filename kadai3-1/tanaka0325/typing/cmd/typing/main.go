package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gopherdojo/dojo8/kadai3-1/tanaka0325/typing"
)

var (
	time     int
	filepath string
)

const (
	timeUsageText     = "game time (second)"
	filepathUsageText = "words file path"
)

func init() {
	flag.IntVar(&time, "time", 10, timeUsageText)
	flag.IntVar(&time, "t", 10, timeUsageText+" (short)")
	flag.StringVar(&filepath, "file", "official/holmes.txt", filepathUsageText)
	flag.StringVar(&filepath, "f", "official/holmes.txt", filepathUsageText+" (short)")
	flag.Parse()
}

func main() {
	ws, err := getWords(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// validate time
	if time <= 0 {
		err := fmt.Errorf("time should be bigger than 0")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else if time > 100 {
		err := fmt.Errorf("time should be smaller than 100")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	game := typing.NewGame(ws, time)
	if err := game.Play(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getWords(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	trimed := strings.Trim(string(content), "\n")
	return strings.Split(trimed, " "), nil
}
