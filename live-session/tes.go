package main

import (
	"fmt"
)

func main() {
	mySlices := []int{1, 2, 3, 4, 5}

	for _, value := range mySlices {
		fmt.Printf("index :  , value : %d\n", value)
	}
}
