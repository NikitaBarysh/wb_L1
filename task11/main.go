package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	res := merge(set1, set2)
	fmt.Println(res)
}

func merge(set1, set2 []int) []int {
	resMap := make(map[int]bool) // создаем мапу куда будем класть элементы множеств
	res := make([]int, 0)        // создаем мапу куда будем класть уникальные значения из 2 слайсов

	for _, v := range set1 { // кладем в мапу элементы из первого множества
		resMap[v] = true
	}

	for _, v := range set2 { // кладем в мапу элементы из второго множества
		resMap[v] = true
	}

	for k := range resMap { // берем ключи, которые взяты из двух множеств и уникальны и кладем в слайс
		res = append(res, k)
	}
	return res
}
