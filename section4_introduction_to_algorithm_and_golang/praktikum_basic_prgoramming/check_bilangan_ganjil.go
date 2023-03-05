package main

import (
	"fmt"
)

func main() {
	// declare variables
	var number int64
	// input processing
	fmt.Print("Input The number")
	fmt.Scanln(&number)
	//  chechk whether the number of variables that have been inputted is exhausted
	//  if divided by 2 if os then it is an even number
	if number%2 == 0 {
		fmt.Println(number, "is an even number")
	} else {
		fmt.Println(number, "is an odd number")
	}

}
