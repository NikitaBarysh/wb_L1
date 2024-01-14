package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abCdefAaf"
	fmt.Println(check(s))
}

func check(s string) bool {
	strings.ToLower(s)      // приводим все символы в строке к нижнему регистру
	arr := []rune(s)        // разбиваем на буквы в unicode
	m := make(map[int]bool) // делаем мапу куда будем класть существующие символы
	for _, val := range arr {
		if _, ok := m[int(val)]; ok == true { // если в мапе есть такой символ, то вернем false
			return false
		} else { // если нету, то кладем этот символ
			m[int(val)] = true
		}
	}

	return true
}
