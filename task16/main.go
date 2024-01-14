package main

import "fmt"

func main() {
	arr := []int{34, 11, 234, 3, 124, 132, 23, 54, 76, 2, 16, 34}

	fmt.Println(sort(arr))
}

func sort(nums []int) []int {
	if len(nums) < 2 { // базовый случай
		return nums
	}

	base := nums[0] // значения относительно которого будем решать куда добавлять

	less := make([]int, 0) // для значений меньше выбранного значения выше
	more := make([]int, 0) // для значений больше

	// в цикле решаем куда добавлять значения, влево или вправо
	for i := 1; i < len(nums); i++ {
		if nums[i] <= base {
			less = append(less, nums[i])
		} else {
			more = append(more, nums[i])
		}
	}

	// рекурсивно все добавляем в один срез
	nums = append(sort(less), base)
	nums = append(nums, sort(more)...)

	return nums
}
