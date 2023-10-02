package main

import "fmt"

func refersedString(str string) string {
	tmp := []rune(str)
	for i, j := 0, len(tmp)-1; i < len(tmp)/2; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}

func main() {
	var str string
	fmt.Scanf("%s", &str)

	str = refersedString(str)
	fmt.Println(str)
}
