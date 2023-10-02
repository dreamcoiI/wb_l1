package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("Before a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("After a: %d, b: %d", a, b)
}
