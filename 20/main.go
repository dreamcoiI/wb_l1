package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '—' // Задайте все необходимые разделители здесь
	})

	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}

	separator := " " // Здесь также можно задать нужный разделитель
	return strings.Join(reversedWords, separator)
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
