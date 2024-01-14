package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2) // создаем тикер

	go stop2(ticker)

	time.Sleep(time.Second * 2) // Ждем 2 секунды и останавливаем тикер

	ticker.Stop()

	time.Sleep(time.Second * 3) // Ждем завершения горутины
}

func stop2(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C: // когда тикер остановят, придет инфа в канал и закончим выполнение
			fmt.Println("Stop")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
