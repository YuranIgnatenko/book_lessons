package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	p := fmt.Println
	var chRes = make(chan int, 1)

	process := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			p("Done")
		case <-chRes:
			fmt.Println("get channel ..")
		}
		p("out ..")
		wg.Done()
	}

	recalcChannel := func() {
		x := 0
		for {
			x++
			time.Sleep(time.Second * 1)
			p("counter channel: ", x)
			if x == 5 {
				chRes <- 99
				defer close(chRes)
				wg.Done()
				return
			}
		}
	}

	recalcCancel := func() {
		x := 0
		for {
			x++
			time.Sleep(time.Second * 1)
			p("counter cancel: ", x)
			if x == 7 {
				wg.Done()
				cancel()
			}
		}
	}

	wg.Add(3)

	go process(ctx)
	go recalcChannel()
	go recalcCancel()

	wg.Wait()

}
