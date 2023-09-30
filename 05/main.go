package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Println("Write seconds :")
	fmt.Scanf("%d", &n)
	ch := make(chan int)

	go Write(ch)

	go Read(ch)

	time.Sleep(time.Duration(n) * time.Second)

	close(ch)

}

func Write(ch chan int) {
	defer close(ch)

	for i := 1; ; i++ {
		fmt.Println("Write value in channel", i)
		ch <- i
		time.Sleep(time.Millisecond * time.Duration(500))
	}
}

func Read(ch chan int) {
	for value := range ch {
		fmt.Println("Read value in chanel:", value)
	}
}
