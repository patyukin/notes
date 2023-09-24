package main

import "fmt"

func integers() chan int {
	ch := make(chan int)

	go func() {
		for i := 1; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	ch := integers()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
