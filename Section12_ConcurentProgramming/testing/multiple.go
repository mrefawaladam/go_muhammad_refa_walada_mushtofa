package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateArray(length int) []int {
	array := make([]int, length)
	rand.Seed(time.Now().UnixNano())

	for i := range array {
		array[i] = rand.Intn(100)
	}

	return array
}

func sumArray(arr []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var sum int
	for _, num := range arr {
		sum += num
	}

	resultChan <- sum

}

func main() {
	array := generateArray((1000))
	resultChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go sumArray(array[i*100:(i+1)*100], resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	fmt.Println(resultChan)
	var sum int
	for num := range resultChan {
		sum += num
	}

	fmt.Printf("Total: %d\n", sum)
}
