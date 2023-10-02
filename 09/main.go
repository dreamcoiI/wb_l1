package main

import (
	"fmt"
	"sync"
)

func main() {
	inputCh := make(chan int)
	outputCh := make(chan int)

	numbers := [10]int{}
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range inputCh {
			result := x * 2
			outputCh <- result
		}
		close(outputCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, x := range numbers {
			inputCh <- x
		}
		close(inputCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range outputCh {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
