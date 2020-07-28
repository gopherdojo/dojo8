package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopherdojo/dojo8/kadai4/tsuchinaga/clock"

	"github.com/gopherdojo/dojo8/kadai4/tsuchinaga/server"
)

func main() {
	serv := server.New(clock.New())
	go func() {
		if err := serv.Run(); err != nil {
			log.Println(err)
			os.Exit(2)
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)
		if serv.GetAddr() != "" {
			fmt.Printf("server started at %s\n", serv.GetAddr())
			break
		}
	}
	<-make(chan error)
}
