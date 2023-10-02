package main

import (
	"fmt"
	"strings"
)

func Unique(str string) bool {
	lcStr := strings.ToLower(str)

	uniqueStr := make(map[rune]bool)

	for _, char := range lcStr {
		if uniqueStr[char] {
			return false // Найден повторяющийся символ
		}
		uniqueStr[char] = true
	}

	return true

}

func main() {
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range stringsToCheck {
		result := Unique(str)
		fmt.Printf("%s — %t\n", str, result)
	}
}
