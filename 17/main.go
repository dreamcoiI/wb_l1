package main

import (
	"fmt"
	"sort"
)

func myBinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		middle := low + (high-low)/2
		if target < arr[middle] {
			high = middle - 1
		} else if target > arr[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 4, 5, 7, 8, 9}

	target := 4

	id := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if id < len(arr) && arr[id] == target {
		fmt.Printf("number %d at index %d\n", target, id)
	} else {
		fmt.Println("Cant find", target)
	}

	id = myBinarySearch(arr, target)
	if id >= 0 {
		fmt.Printf("number %d at index %d\n", target, id)
	} else {
		fmt.Println("Cant find", target)
	}
}
