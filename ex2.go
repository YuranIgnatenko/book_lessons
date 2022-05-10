package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const MAX_ITER_RECALC = 10

func clear() {
	fmt.Println("\033[H\033[2J")
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	p := fmt.Println
	pf := fmt.Printf
	var chRes = make(chan int, 1)
	defer close(chRes)

	process := func(ctx context.Context) {
		p("START GET SIGNALS")
		select {
		case <-ctx.Done():
			p("get ctx.Done ..")
			chRes <- 99
			cancel()
			wg.Done()
		}
	}

	recalcCancel := func() {
		p("START RECALC")
		percent := 0
		for {
			if chRes != nil {
				wg.Done()
				cancel()
			}
			percent++
			time.Sleep(time.Millisecond * 500)
			pf("counter cancel: [%+v%s]\n", percent, "%")
			if percent == MAX_ITER_RECALC {
				wg.Done()
				cancel()
			}
		}
	}

	stopperFromTimerChannel := func() {
		time.Sleep(time.Second * 2)
		chRes <- 99
		wg.Done()
	}

	// stopperFromTimerCtx := func() {
	// 	time.Sleep(time.Second * 4)
	// 	wg.Done()
	// 	cancel()
	// }

	wg.Add(3)

	go process(ctx)
	go recalcCancel()
	go stopperFromTimerChannel()
	// go stopperFromTimerCtx()
	wg.Wait()

}
