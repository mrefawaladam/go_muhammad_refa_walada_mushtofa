package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The recipient goroutine is required to receive data must be prepared to receive the data
func printNumberUnbuffered(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func printNumberBuffered(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func printRandomNumber() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}

func main() {
	go printRandomNumber()
	time.Sleep(1 * time.Second)
	// unbuffered
	ch := make(chan int)
	go printNumberUnbuffered(ch)
	for num := range ch {
		fmt.Println(num)
	}
	// buffered
	chb := make(chan int, 5)
	go printNumberBuffered(chb)
	for num := range chb {
		fmt.Println(num)
	}

}
