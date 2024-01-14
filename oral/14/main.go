package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}

/*
Ответ: [b b a][a a]
Мы передаем в анонимную функцию слайс с емкостью 2, так как мы сразу добавляем элемент, то увеличивается емкость и
выделяется новый участок памяти для этого слайса и все изменения будут только в новом слайсе, а старый не изменится
*/
