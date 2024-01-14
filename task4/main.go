package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var workers int // объявляем переменную, куда будем записывать воркеров
	fmt.Println("Введите кол-во воркеров")
	fmt.Scanln(&workers) // ждем когда введут кол-во воркеров

	jobCh := make(chan int) // создаем не буферизированный

	termSing := make(chan os.Signal, 1)                      // создаем канал, в который будем ждать сигнал о завершении программы
	signal.Notify(termSing, syscall.SIGTERM, syscall.SIGINT) //какие сигналы ждем

	for i := 0; i < workers; i++ { // создаем воркеров
		go worker(jobCh) // запускаем в горутине воркера с каналом куда записываем данные
	}

	go func() { // запускаем в горутине аноним функцию и вней делаем беск цикл, где записываем в канал рандом значения
		for {
			jobCh <- rand.Int()
			time.Sleep(time.Second * 1)
		}
	}()

	fmt.Println(<-termSing) // выводим сигнал о завершении

}

func worker(job <-chan int) {
	for {
		select {
		case data, ok := <-job: // получаем данные из канала и проверяем не закрыт ли канал
			if !ok {
				fmt.Println("Канал закрыт")
			}
			fmt.Println(data) // выводим данные из канала
		}
	}
}
