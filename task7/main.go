package main

import (
	"fmt"
	"sync"
)

func main() {
	dict := make(map[int]int) // объявляем мапу
	var wg sync.WaitGroup     // для ожидания выполнения всех горутин
	var mu sync.Mutex         // для конкурентной записи в мапу

	for i := 0; i < 10; i++ {
		wg.Add(1) // добавляем задачу
		go func(i int) {
			defer wg.Done() // говорим, что выполнили задачу
			// блокируем запись для других горутин, они выстраиваются в очередь, после записи, разблокируем доступ
			// и все повторяется с другими горутинами
			mu.Lock()
			defer mu.Unlock()
			dict[i] = i
			fmt.Println(dict[i])

		}(i)
	}
	wg.Wait() // ждем выполнения всех задач

	fmt.Println("Final:", dict)
}
