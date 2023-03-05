package main

import (
	"fmt"
)

func fibonacci(number int) int {
	// check number is not 1/0
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		// cehck  number is greater than 1, recursively call the fibonacci function with arguments number-1 and number-2, and return the sum.
		return fibonacci(number-1) + fibonacci(number-2)
	}
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144

}
