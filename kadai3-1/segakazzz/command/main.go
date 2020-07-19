package main

import (
	"flag"
	"github.com/gopherdojo/dojo8/kadai3-1/segakazzz/tpgame"
	"os"
	"time"
)

func main() {
	var (
		json string
		sec int
	)
	flag.StringVar(&json, "json", "../testdata/words.json", "Path to source json file to import")
	flag.IntVar(&sec, "sec", 10, "Seconds to timeout")

	err := tpgame.Run(json, time.Duration(sec) * time.Second)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
