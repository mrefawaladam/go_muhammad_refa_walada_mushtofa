package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World!") // print text

	// variable in golang

	var number int8 = 1
	fmt.Println(number)

	var student1 string = "Refa" //type is string : in goalang you can specify the data type
	var student2 = "Beril"       // type is inferred : accpets all types without being defined
	x := "Refa Ganteng"          // type is inferred : accpets all types without being defined

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var nameNull string // hwo to write empty variable in golang
	fmt.Println(nameNull)

	const phi = 3.14 // how to write const in golang
	const appName = "My App"

	// Oprator in golang

	// Oprator Aritmatika
	// a := 18
	// b := 5

	// c := (a == b) // false
	// d := (a >  b) // true
	// e := (a <= b) // false

	// Operator Bitwise
	// a := 5 // representation bit: 0101
	// b := 3 // representation bit: 0011

	// c := a & b // return 1 (0001)
	// d := a | b // return 7 (0111)
	// e := a ^ b // return 6 (0110)
	// f := a << 1 // return 10 (1010)
	// g := a >> 1 // return 2 (0010)

}
