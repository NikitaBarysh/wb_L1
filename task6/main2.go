package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст с возможностью отмены
	go stop1(ctx)                                           //
	time.Sleep(time.Second * 3)                             // Ждем 3 секунды и отменяем выполнение

	cancel() //

	time.Sleep(time.Second * 1) // Ждем завершения горутины
}

func stop1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // когда вызовем метод  cancel(), придет в канал сигнал и остановим выполнение
			fmt.Println("Stop")
			time.Sleep(time.Second * 2)
			return
		default:
			fmt.Println("Working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
