package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go stop3(ch)

	time.Sleep(time.Second * 2) // Ждем 2 секунды и закрываем канал

	close(ch)

	time.Sleep(time.Second * 2) // Ждем завершения горутины
}

func stop3(ch chan struct{}) {
	for {
		fmt.Println("Working")

		_, ok := <-ch // проверяем на закрытие канала
		if !ok {      // если закрыт, то заканчиваем выполнение
			fmt.Println("Stop")
			return
		}

	}
}
