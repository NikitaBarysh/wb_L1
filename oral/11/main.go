package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

/*
4
2
3
0
1
fatal error: all goroutines are asleep - deadlock!

Надо передавать указатель WaitGroup, команда wg.Done выполняется локально и, когда дойдем до wg.Wait(),
main будет остановлена
*/
