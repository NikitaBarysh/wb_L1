package main

import "fmt"

func main() {
	var number int64 = 11

	fmt.Printf("Исходное значение: %d\n", number)

	// Устанавливаем 3-й бит в 1
	setBitTo1(&number, 3)
	fmt.Printf("Установлен 3-й бит в 1: %d\n", number)

	// Устанавливаем 2-й бит в 0
	setBitTo0(&number, 2)
	fmt.Printf("Установлен 2-й бит в 0: %d\n", number)
}

func setBitTo1(num *int64, i int) {
	// Устанавливаем i-й бит в 1
	*num |= 1 << i
}

func setBitTo0(num *int64, i int) {
	// Устанавливаем i-й бит в 0
	*num &^= 1 << i
}
