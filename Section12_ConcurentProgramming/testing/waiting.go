package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("Hello")
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("World")
		}
	}()

	wg.Wait()
	fmt.Println("Done")
}
