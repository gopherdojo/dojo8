package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo8/kadai4/segakazzz/omikuji"
	"os"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "Port Number to serve http")
}

func main() {
	flag.Parse()
	err := omikuji.Run(port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
