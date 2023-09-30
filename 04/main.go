package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	wg *sync.WaitGroup
	ch chan int
}

type PoolWorker struct {
	Pool []Worker
}

func PoolCreate(n int, wg *sync.WaitGroup, ch chan int) *PoolWorker {
	pW := make([]Worker, 0, n)
	for i := 0; i < n; i++ {
		pW = append(pW, Worker{wg: wg, ch: ch})
	}
	result := PoolWorker{Pool: pW}
	return &result
}

func (w *Worker) Work(id int) {
	defer w.wg.Done()
	for data := range w.ch {
		fmt.Printf("Worker#%d got data - %d\n", id, data)
	}
}

func (p *PoolWorker) WorkStart() {
	for i, worker := range p.Pool {
		go worker.Work(i)
	}
}

func Shutdown(ch chan int, wg *sync.WaitGroup) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	close(ch)

	wg.Wait()
}

func main() {

	var n int
	fmt.Println("Number workers: ")
	fmt.Scanf("%d", &n)

	var wg sync.WaitGroup

	wg.Add(n)
	ch := make(chan int)

	pool := PoolCreate(n, &wg, ch)

	pool.WorkStart()

	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	Shutdown(ch, &wg)

}
