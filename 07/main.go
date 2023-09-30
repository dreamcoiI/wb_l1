package main

import (
	"fmt"
	"sync"
)

func main() {
	data := new(sync.Map)

	numWorkers := 5
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		key := fmt.Sprintf("key%d", i)
		value := i * i
		wg.Add(1)

		go func(key string, value int) {
			defer wg.Done()
			data.Store(key, value)

			fmt.Printf("Worker write in map: %s -> %d\n", key, value)
		}(key, value)
	}
	wg.Wait()

	fmt.Println("Содержимое map после всех записей:")
	data.Range(func(key, value interface{}) bool {
		fmt.Printf("%s -> %v\n", key, value)
		return true
	})
}
