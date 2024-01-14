package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ { // создаем цикл, где будем запускать 5 горутин
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ { // цикл на запись 100 значений в counter
				mu.Lock() // ставим блокировку, чтобы в counter записывала только одна горутина

				counter++

				mu.Unlock() // снимаем блокировку после прибавки
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
