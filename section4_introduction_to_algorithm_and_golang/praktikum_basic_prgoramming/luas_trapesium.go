package main
import (
	"fmt"
)

func main() {

	// declaration variables
	var top, bottom, height float32
	// create a function to perform the input process for top bottom and height variables
	fmt.Print("Enter the length of the top side of the trapesium : ")
	fmt.Scanln(&top) // Tanda & digunakan untuk mengambil alamat variabel teratas

	fmt.Print("Input the length of the bottom side of the trapesium : ")
	fmt.Scanln(&bottom)

	fmt.Print("Input height of the trapesium : ")
	fmt.Scanln(&height)

	// create an wide variable function to do the calculation that holds from the formula  to find the area of a trapesim
	wide := 0.5 * (top + bottom)* top

	fmt.Printf("The area of the trapesium is : %.2f\n",wide)
}