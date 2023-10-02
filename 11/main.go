package main

import "fmt"

func main() {
	set1 := []int{1, 3, 5, 7, 9}
	set2 := []int{5, 1, 10, 123, 12, 3}

	intersect := make([]int, 0)
	set2Map := make(map[int]bool)

	for _, num := range set2 {
		set2Map[num] = true
	}

	for _, num := range set1 {
		if set2Map[num] {
			intersect = append(intersect, num)
		}
	}
	fmt.Println(intersect)
}
