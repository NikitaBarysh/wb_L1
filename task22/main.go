package main

import (
	"fmt"
	"math/big"
)

func main() {

	// объявляем переменные
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)

	// присваиваем значение
	a.SetString("4325625642365231654575698245762147", 10)
	b.SetString("2562457654765368245652678452667454", 10)

	// Делаем операции
	fmt.Println("Сложение:", res.Add(a, b))
	fmt.Println("Вычитание:", res.Sub(a, b))
	fmt.Println("Умножение:", res.Mul(a, b))
	fmt.Println("Деление:", res.Div(a, b))

}
