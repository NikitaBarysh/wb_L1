package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	i := 2 // индекс
	out := remove(nums, i)

	fmt.Println(out)
}

func remove(nums []int, idx int) []int {
	// проверяем индекс(не отрицательный ли он и не выходит ли он за пределы слайса)
	if idx < 0 || idx >= len(nums) {
		return nums
	}
	//добавляем в слайс все что до индексов и после, тем самым удаляя нужный элемент
	nums = append(nums[:idx], nums[idx+1:]...)
	return append(nums)
}
