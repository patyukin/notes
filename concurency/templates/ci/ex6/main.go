package main

import "fmt"

// sendItems отправляет элементы в канал
func sendItems(ch chan<- int) {
	// Закрываем канал после отправки всех элементов
	defer close(ch)

	for i := 1; i <= 5; i++ {
		ch <- i // Отправляем элементы в канал
	}
}

func main() {
	ch := make(chan int) // Создаем канал

	go sendItems(ch) // Запускаем функцию отправки элементов в канал в отдельной горутине

	// Итерация по элементам, получаемым из канала
	for item := range ch {
		fmt.Println(item)
	}
}
