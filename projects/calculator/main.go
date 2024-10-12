package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		select {
		case x := <-firstChan:
			out <- x * x
		case x := <-secondChan:
			out <- x * 3
		case <-stopChan:
			return
		}
	}()

	return out
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	resultChan := calculator(firstChan, secondChan, stopChan)

	// Пример 1: Отправка значения в первый канал (квадрат числа)
	go func() {
		firstChan <- 4
	}()

	// Пример 2: Отправка значения во второй канал (умножение на 3)
	//go func() {
	//	secondChan <- 5
	//}()

	// Пример 3: Завершение работы через канал stopChan
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	close(stopChan)
	// }()

	select {
	case result := <-resultChan:
		fmt.Println("Результат:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Таймаут!")
	}
}
