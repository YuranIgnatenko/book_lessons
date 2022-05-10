package main

import (
	"os/exec"
	"sync"
	"context"
	"time"
	"fmt"
)

func main(){
	// RunChan()
	// RunContext()
	// RunBuffChan()
	// RunSync1()
	// RunSync2()
	// RunBinary("./proc1.sh")
	// AbortBinary()
	// AbortBinary2()
}


func AbortBinary2(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	data := make([]int,0)
	go func(){
		fmt.Scanln()
		data = make([]int,0)
		cancel()
	}()
	go func(){
		for i:=0;i<=100;i++{
			fmt.Print("\033[H\033[2J")
			fmt.Printf("process ---- [%d%s]", i, "%")
			data = append(data, i)
			time.Sleep(time.Millisecond * 50)
		}
		cancel()
	}()

	fcheck := func(ctx context.Context){
		select{
		case <- time.After(time.Second * 50):
			fmt.Println("\nDone Process !")
		case <- ctx.Done():
			fmt.Println("\nCanceled Process")
		}
		fmt.Println(data)
	}
	fcheck(ctx)
}

func AbortBinary(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func(){fmt.Scanln();cancel()}()
	go func(){
		for i:=0;i<=100;i++{
			fmt.Print("\033[H\033[2J")
			fmt.Printf("process ---- [%d%s]", i, "%")
			time.Sleep(time.Millisecond * 50)
		}
		cancel()
	}()

	fcheck := func(ctx context.Context){
		select{
		case <- time.After(time.Second * 50):
			fmt.Println("\nDone Process !")
		case <- ctx.Done():
			fmt.Println("\nCanceled Process")
		}
	}
	fcheck(ctx)
}

func RunBinary(name string){
	exec.Command(name, "&")
}

func RunSync2(){
	var wg sync.WaitGroup
	f1:=func(){fmt.Println("f1");wg.Done()}
	f2:=func(){fmt.Println("f2");wg.Done()}
	f3:=func(){fmt.Println("f3");wg.Done()}
	f4:=func(){fmt.Println("f4");wg.Done()}
	f5:=func(){fmt.Println("f5");wg.Done()}
	ar_f := []func(){f1,f2,f3,f4,f5}

	for _, function := range ar_f{
		wg.Add(1)
		function()
	}
	wg.Wait()
}

func RunSync1(){
	data := []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}

	var wg sync.WaitGroup

	proc := func(dt string){
		fmt.Println(dt)
		wg.Done()
	}

	for _, elem := range data{
		wg.Add(1)
		go proc(elem)
	}

	wg.Wait()

}

func RunBuffChan(){
	// if size buffer < 2, then deadlock
	msg := make(chan string, 2)
	msg <- "test"
	fmt.Println(<- msg)
}

func RunContext(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func(){fmt.Scanln();cancel()}()

	process := func(ctx context.Context){
		select{
		case <- time.After(time.Second * 4):
			fmt.Println("Done process")
		case <- ctx.Done():
			fmt.Println("Canceled process")
		}
	}
	process(ctx)
}

func RunChan(){
	m1 := make(chan string)
	m2 := make(chan string)

	go func(){
		for {
			time.Sleep(time.Second * 2)
			m1 <- "go func 1"
		}
	}()
	go func(){
		for {
			time.Sleep(time.Millisecond * 500)
			m2 <- "go func 2"
		}
	}()

	for{
		select{
		case m := <- m1:
			fmt.Println("case 1", m)
		case m := <- m2:
			fmt.Println("case 2", m)
		}
	}
}