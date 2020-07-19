package main

import (
	"github.com/gopherdojo/dojo8/kadai3-1/segakazzz/tpgame"
	"os"
	"time"
)

func main (){

	err := tpgame.Run("../testdata/words.json", 10 * time.Second)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

