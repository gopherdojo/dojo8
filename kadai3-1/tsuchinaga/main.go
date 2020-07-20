package main

import (
	"bufio"
	"os"
	"time"

	"github.com/gopherdojo/dojo8/kadai3-1/tsuchinaga/typinggame"
)

func main() {
	bs := bufio.NewScanner(os.Stdin)
	bs.Split(bufio.ScanLines)

	g := typinggame.New(time.Now().UnixNano())

	tCh := time.After(15 * time.Second)
	for {
		iCh := make(chan string)
		go func() {
			println(g.Next())
			bs.Scan()
			iCh <- bs.Text()
		}()

		select {
		case ans := <-iCh:
			g.Answer(ans)
		case _ = <-tCh:
			println("\nResult", g.Result())
			return
		}
	}
}
