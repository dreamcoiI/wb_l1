package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
}

func main() {
	fmt.Println("Start")
	Sleep(4)
	fmt.Println("END")
}
