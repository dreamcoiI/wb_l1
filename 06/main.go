package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	doneChan  = make(chan struct{})
	ctx       context.Context
	cancelCtx context.CancelFunc
	flag      = false
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go withChannel(&wg)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("Goroutine stopped with channel")
	doneChan <- struct{}{}

	wg.Add(1)

	ctx, cancelCtx = context.WithCancel(context.Background())
	go withContext(&wg)
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("Goroutine stopped with channel")
	cancelCtx()

	wg.Add(1)
	go withFlag(&wg)
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("Goroutine stopped with channel")
	flag = true

	wg.Wait()
}

func withChannel(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-doneChan:
			fmt.Println("Goroutine with changelist is stopped ")
			return
		default:
			fmt.Println("Goroutine with changelist is working")
			time.Sleep(time.Second)
		}
	}
}

func withContext(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine with context is stopped ")
			return
		default:
			fmt.Println("Goroutine with context is working")
			time.Sleep(time.Second)
		}
	}
}

func withFlag(wg *sync.WaitGroup) {
	defer wg.Done()
	for !flag {
		fmt.Println("Goroutine with flag is stopped ")
		return
	}
	fmt.Println("Goroutine with flag is working")
	time.Sleep(time.Second)

}
