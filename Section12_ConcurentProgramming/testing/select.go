package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "Hello"
	}()

	go func() {
		chan2 <- "world"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
