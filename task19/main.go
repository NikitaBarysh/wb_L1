package main

import "fmt"

func main() {
	s := "главрыба"
	fmt.Println(rev(s))
}

func rev(s string) string {
	res := []rune(s) // делаем срез рун из строки

	l := 0            // первый символ строки
	r := len(res) - 1 // последний символ строки
	for l < r {
		res[l], res[r] = res[r], res[l] // меняем местами буквы
		r--                             // двигаемся с конца влево
		l++                             // двигаемся с начала вправо
	}

	return string(res) // возвращаем перевернутую строку
}
