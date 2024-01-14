package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun "
	fmt.Printf("%s - %s", s, rev(s))

}

func rev(s string) string {
	res := strings.Split(s, " ") // разделяем строку на слова

	l := 0
	r := len(res) - 1

	for l < r { // меняем местами
		res[l], res[r] = res[r], res[l]
		r--
		l++
	}

	return strings.Join(res, " ") // соединяем обратно в строку
}
