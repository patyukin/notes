package main

import (
	"fmt"
)

// Generator - функция, которая генерирует значения в горутине и отправляет их по каналу
func Generator(done chan bool) chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		done <- true
		close(out)
	}()

	return out
}

func main() {
	// канал для определения окончания генерации
	done := make(chan bool)

	c := Generator(done)

	// запускаем в горутине, чтобы избежать блокировки
	go func() {
		<-done
	}()

	// итерация по каналу
	for i := range c {
		fmt.Println(i)
	}
}
