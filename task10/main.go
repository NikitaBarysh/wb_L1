package main

import (
	"fmt"
)

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := groupTemp(data)
	fmt.Println(res)
}

func groupTemp(temp []float64) map[int][]float64 {
	group := make(map[int][]float64) // создаем мапу куда будет группировать

	for _, v := range temp { //обходим слайс со значениями
		key := int(v) / 10 * 10            // находим ключ группы
		group[key] = append(group[key], v) // добавляем в нужную группу
	}

	return group
}
