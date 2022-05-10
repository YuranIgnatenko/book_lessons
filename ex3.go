package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	p := fmt.Println

	recalc := func() {
		for i := 0; i <= 100000; i++ {
			if ctx.Done() == nil {
				p("exit ..")
				os.Exit(0)
			}
			p("[", i, "]")
		}
	}

	hook := func() {
		time.Sleep(time.Second * 2)
		p("cancel ..")
		cancel()
		wg.Done()
	}
	wg.Add(1)
	go hook()
	recalc()
	wg.Wait()
}
