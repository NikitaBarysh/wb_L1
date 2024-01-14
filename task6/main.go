package main

import (
	"fmt"
	"time"
)

func main() {
	doneCh := make(chan struct{})

	go stop(doneCh)

	time.Sleep(time.Second * 3) // Ждем 3 секунды и отменяем выполнение

	doneCh <- struct{}{} // пишем в канал пустую структуру

	time.Sleep(time.Second * 1) // Ждем завершения горутины

}

func stop(doneCh chan struct{}) {
	for {
		select {
		case <-doneCh: // когда прейдет пустая структура, завершим выполнение
			fmt.Println("Stop")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Second)
		}
	}
}
