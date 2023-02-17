package main

import (
	"fmt"
)

func main() {

	// print text
	// fmt.Println("Hello World!")

	var number int8 = 1
	fmt.Println(number)

	var student1 string = "Refa" //type is string : in goalang you can specify the data type
	var student2 = "Beril"       // type is inferred : accpets all types without being defined
	x := "Refa Ganteng"          // type is inferred : accpets all types without being defined

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

}
