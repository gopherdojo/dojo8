package main

import (
	"math/rand"
	"time"

	"github.com/gopherdojo/dojo8/kadai4/hokita/omikuji/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	server.Run()
}
