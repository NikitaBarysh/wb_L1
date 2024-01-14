package main

import "fmt"

func main() {
	input := []int{2, 4, 6, 8, 10}

	doneCh := make(chan struct{}) // канал для сигнала выхода из горутины
	defer close(doneCh)           // при завершении программы, закрываем канал, чтобы все горутины завершились

	inputCh := generator(doneCh, input) // из слайса в канал

	results := multiply(doneCh, inputCh) //берем значения из канала выше, умножаем, кладем во второй канал

	for res := range results { //выводим значения из второго канала
		fmt.Println(res)
	}
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	multiplyRes := make(chan int) // создаем канал куда будем класть данные, которые умножили на 2

	go func() {
		defer close(multiplyRes) // когда завершаем горутину, закрываем канал

		for data := range inputCh { //берем значения из первого канала
			res := data * 2 // умножаем их
			select {
			case <-doneCh: // когда закроется канал выходим из горутины
				return
			case multiplyRes <- res: // кладем умноженные значения в канал
			}
		}
	}()

	return multiplyRes // возвращаем канал с умноженными значениями
}

func generator(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int) // канал куда будем класть значения из слайса

	go func() {
		defer close(inputCh) // закрываем канал, при завершении горутины

		for _, data := range input {
			select {
			case <-doneCh: // когда закроется канал выходим из горутины
				return
			case inputCh <- data: // кладем значения из слайса  в канал
			}
		}
	}()

	return inputCh // // возвращаем канал со значениями из слайса
}
