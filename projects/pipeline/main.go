package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	var prev string
	first := true

	for s := range inputStream {
		if first || s != prev {
			outputStream <- s
			prev = s
			first = false
		}
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	// Запускаем функцию removeDuplicates в отдельной горутине
	go removeDuplicates(inputStream, outputStream)

	// Отправляем данные в inputStream
	go func() {
		values := []string{"apple", "apple", "banana", "banana", "apple", "pear", "pear", "banana"}
		for _, v := range values {
			inputStream <- v
		}
		close(inputStream) // Закрываем канал, чтобы завершить цикл в removeDuplicates
	}()

	// Читаем данные из outputStream и выводим их
	for v := range outputStream {
		fmt.Println(v)
	}
}
