package main

import (
	"fmt"
)

func main() {
	price := map[string]int{"siomay":40000,"batagor":3000}
	fmt.Println(price)

	// simpel
	var priceDiscount = make(map[string]int)
	priceDiscount["bag"]= 3999
	fmt.Println(priceDiscount)
}

