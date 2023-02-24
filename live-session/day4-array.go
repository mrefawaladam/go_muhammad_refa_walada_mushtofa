package main

import (
	"fmt"
)

func main() {

	// var cities = []string
	// cities := [100]string{"Bandung"}
	// propertties
	// pinter

	cities := ([]string{"Bandung", "Bandung"})
	fmt.Println(cities)
	
	citiesInt := make([]int, 5,7)
	fmt.Println(citiesInt)
	citiesInt = append(citiesInt,100)
	fmt.Println(len(citiesInt) , cap(citiesInt))
	fmt.Println(citiesInt)

	// convert int 
	var i32 int32 = 34
	var i64 int64 = int64(i32) 
	fmt.Printf("%T", i64)

	// map 
	var nilaiSiswa map[string]int32
	
}	