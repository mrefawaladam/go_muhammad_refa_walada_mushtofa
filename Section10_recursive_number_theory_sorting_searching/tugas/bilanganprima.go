package main

import "fmt"

func primeX(number int) int {
	count := 0
	// start looping from 2, which is the smallest prime number
	for i := 2; ; i++ {
		prime := true
		// loop throung 2 to the square root of i to chechk if number id devisible by any numbers
		for j := 2; j*j <= i; j++ {
			// perform the determination to determine prime numbers
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			count++
			if count == number {
				return i
			}
		}
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
