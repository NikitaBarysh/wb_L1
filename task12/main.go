package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := uniq(data)
	fmt.Println(set)
}

func uniq(set []string) []string {
	res := make(map[string]bool) // создаем мапу куда будем класть значения из последовательности
	out := make([]string, 0)     // множество

	// кладем последовательность и кладем туда значения из последовательности и когда будет повторяться мы будем как бы
	// перезаписывать по ключу значение и таким образом мы сделаем уникальность
	for _, v := range set {
		res[v] = true
	}

	// обходим мапу и достаем ключ, который стал уникальным значением
	for k := range res {
		out = append(out, k) // кладем в множество этот ключ
	}

	return out
}
