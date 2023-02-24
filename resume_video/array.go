package main 

import (
	"fmt"
	"reflect"
)
// pull request git
func main()  {
	var primes [5]int
	var couteries [5]string

	fmt.Println(reflect.ValueOf(asdfasd).Kind())
	fmt.Println(reflect.ValueOf(adfasf).Kind())

	// how to write arrat with default value
	x := [5]int {1, 2, 3, 4,5}
	var y [5]int = [5]int{1,2,4}

	fmt.Println(x)
	fmt.Println(y)

	// how to show elements in array
	for index, element := range x {
		fmt.Println(index, "=>",element)	
	}

}