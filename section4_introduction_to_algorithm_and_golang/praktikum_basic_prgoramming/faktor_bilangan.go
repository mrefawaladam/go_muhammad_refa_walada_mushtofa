package main

import (
	"fmt"
)

func main() {

    var number int
    fmt.Print("Enter a number: ")
    fmt.Scanln(&number)

    fmt.Printf("Factor of a number %d is: ", number)
    for i := 1; i <= number; i++ {
        if number%i == 0 {
            fmt.Printf("%d ", i)
            fmt.Println()
        }
       
    }
}