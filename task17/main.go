package main

import "fmt"

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("значение под индексом:", binarySearch(sorted, 3))
}

func binarySearch(nums []int, target int) int {
	min := 0             // крайний левый индекс
	max := len(nums) - 1 // крайний правый индекс

	for min <= max {
		mid := (max + min) / 2 // середина
		guess := nums[mid]     // возможное значение

		if target == guess { // если равны, то возвращаем
			return mid
		} else if guess > target { // если больше нужного, то убираем правую часть
			max = mid - 1
		} else { // если меньше нужного, то убираем левую часть
			min = mid + 1
		}
	}
	return min // возвращаем индекс, где должен стоять
}
