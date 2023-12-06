package main

import (
	"fmt"
)

func ChannelIterator(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for value := range in {
			out <- value
		}
	}()

	return out
}

func main() {
	in := make(chan int)
	go func() {
		defer close(in)

		for i := 1; i <= 5; i++ {
			in <- i
		}
	}()

	iterator := ChannelIterator(in)
	for value := range iterator {
		fmt.Println(value)
	}
}
