package main

import (
	"fmt"
)

// CreateNumberGenerator создается канал для трансляции чисел от 0 до 9
// и возвращаем этот канал для последующего использования.
func CreateNumberGenerator() <-chan int {
	numChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	return numChan
}

func main() {
	// мы вызываем CreateNumberGenerator(), результат (канал с числами)
	// присваиваем переменной iterator
	iterator := CreateNumberGenerator()

	// итерация по значениям, получаемым из канала.
	// Цикл завершится, как только канал будет закрыт внутри функции CreateNumberGenerator
	for num := range iterator {
		fmt.Println(num)
	}
}
