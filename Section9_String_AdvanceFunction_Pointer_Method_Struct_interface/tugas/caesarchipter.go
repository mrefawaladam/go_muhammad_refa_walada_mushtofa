package main

import "fmt"

func caesar(offset int, input string) string {
	// convert the input string to a slice of runes
	runes := []rune(input)
	// loop through the slice of runes, shifting each rune by the given offset
	for i, r := range runes {
		// check if the current rune is not a space
		if r != ' ' {
			// calculate the shift value
			shift := rune(offset % 26)
			// if the shifted rune goes beyond z
			if r+shift > 122 {
				// shift the rune by wrapping around to a
				runes[i] = rune(96 + int((r+shift)%122))
			} else {
				runes[i] = r + shift
			}
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cngc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
