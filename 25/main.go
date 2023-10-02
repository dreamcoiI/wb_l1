package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {

	<-time.After(duration * time.Second)

}

func main() {

	fmt.Println("START")

	Sleep(2)

	fmt.Println("END")

}
