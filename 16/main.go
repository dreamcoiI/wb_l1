package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, equal, greater []int

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	result := append(less, equal...)
	result = append(result, greater...)
	return result
}

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}

	sortedArr := quickSort(arr)

	fmt.Println(sortedArr)
}
