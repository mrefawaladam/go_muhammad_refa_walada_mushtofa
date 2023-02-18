package main

import (
	"fmt"
)

func main() {
	// declare variables
	var scores int64
	// input processing
	fmt.Print("Input Score: ")
	fmt.Scanln(&scores)
	// i chose the swich case for measurement handling to determine the score
	//  because it look neater thang using an if else
	switch {
	case scores < 0 || scores > 100:    // i am hendling to resolve values less than 0 and more than 100 by using this function
		fmt.Println("Score Invalid")
	case scores >= 80 && scores <= 100:
		fmt.Println("Grade : A")
	case scores >= 65 && scores <= 79:
		fmt.Println("Grade : B")
	case scores >= 50 && scores <= 64:
		fmt.Println("Grade : C")
	case scores >= 35 && scores <= 49:
		fmt.Println("Grade : D")
	default:
		fmt.Println("Grade : E")
	}

}