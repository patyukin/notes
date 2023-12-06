// readme.md
package main

import (
	"fmt"
	"sync"
)

func GenerateData(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i // Отправляем значения в канал
	}
	close(ch) // Закрываем канал после отправки всех данных
}

func ProcessData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик ожидания группы sync.WaitGroup при завершении функции
	for num := range ch {
		squared := num * num
		fmt.Println(squared) // Выводим квадраты чисел на экран
	}
}

func main() {
	ch := make(chan int)    // Создаем канал
	var wg sync.WaitGroup   // Создаем группу ожидания sync.WaitGroup
	wg.Add(1)               // Увеличиваем счетчик ожидания группы на 1
	go GenerateData(ch)     // Запускаем функцию GenerateData в отдельной горутине для заполнения канала
	go ProcessData(ch, &wg) // Запускаем функцию ProcessData в отдельной горутине для обработки данных из канала
	wg.Wait()               // Ожидаем завершения работы всех горутин в группе sync.WaitGroup
}
