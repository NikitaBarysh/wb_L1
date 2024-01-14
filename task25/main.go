package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	sleep(time.Second * 3)
	fmt.Println("stop")

	fmt.Println("start1")
	sleep1(time.Second * 4)
	fmt.Println("stop1")
}

func sleep(i time.Duration) {
	ticker := time.NewTicker(i) // создаем тикер, который через заданное время посылает сообщение в канал и выходит из
	// функции
	<-ticker.C
}

func sleep1(i time.Duration) {
	<-time.After(i) // посылает сообщение в канал через заданное и время и завершает работу функции
}
