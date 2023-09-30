package main

import (
	"fmt"
)

func setBit(n int64, i uint, bitValue uint) int64 {
	if bitValue == 1 {
		return n | (1 << i)
	} else if bitValue == 0 {
		return n & ^(1 << i)
	} else {
		fmt.Println("Неверное значение bitValue. Используйте 0 или 1.")
		return n
	}
}

func main() {
	var n int64
	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Ошибка ввода числа.")
		return
	}

	var i uint
	fmt.Print("Введите индекс бита для изменения: ")
	_, err = fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Ошибка ввода индекса бита.")
		return
	}

	var bitValue uint
	fmt.Print("Введите новое значение бита (0 или 1): ")
	_, err = fmt.Scanf("%d", &bitValue)
	if err != nil {
		fmt.Println("Ошибка ввода значения бита.")
		return
	}

	result := setBit(n, i, bitValue)
	fmt.Printf("Значение после изменения %d-го бита в %d: %d\n", i, bitValue, result)
}
