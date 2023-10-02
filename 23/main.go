package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	i := 5

	if i > len(arr)-1 {
		fmt.Println("ERROR")
		return
	}
	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr)
}
