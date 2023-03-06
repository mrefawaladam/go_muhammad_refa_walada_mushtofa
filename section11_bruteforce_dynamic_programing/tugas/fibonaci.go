package main

import (
	"fmt"
)

func fibo(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a := 0
	b := 1

	// Do iteration as much as n-1 times
	for i := 2; i <= n; i++ {
		// Calculate the nth Fibonacci value
		c := a + b

		// Move the value of b to a and c to b for the next iteration
		a, b = b, c
	}

	// Returns the nth Fibonacci value
	return b
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}
