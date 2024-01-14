package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var timeSec int // объявляем переменную, куда будем записывать через скок сек завершиться программа
	fmt.Println("Введите через сколько завершится программа:")
	fmt.Scan(&timeSec) // // ждем когда введут секунды
	// создаем контекст, который завершится через определенное время
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeSec))
	defer cancel()

	resCh := make(chan int) // канал для записи и чтения

	go func() { // делаем последовательную запись данных
		count := 1
		for {
			resCh <- count
			count++
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// ждем когда в канал поступят данные и читаем их, когда подойдет время к завершению, придет сигнал в ctx.Done()
	// и выйдем из программы
	for {
		select {
		case v := <-resCh:
			fmt.Println(v)
		case <-ctx.Done():
			return
		}
	}
}
