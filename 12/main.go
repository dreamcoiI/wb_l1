package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, str := range arr {
		set[str] = true
	}

	fmt.Println(set)
}
