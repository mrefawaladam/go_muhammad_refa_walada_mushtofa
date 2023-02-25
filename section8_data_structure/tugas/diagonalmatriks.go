package main

import "fmt"

func diagonalDiff(matrisk [][]int) int {
	var leftToRight, rightToLeft int
	n := len(matrisk)

	// loop to count the diagonal number from left to right and from right to left
	for i := 0; i < n; i++ {
		// Number of diagonal elements from left to right
		leftToRight += matrisk[i][i]
		// number of diagonal elements from right to left
		rightToLeft += matrisk[i][n-i-1]
	}
	// restore the difference between the two diagonals
	result := (leftToRight - rightToLeft)

	if result < 0 {
		return -result
	}
	return result

}

func main() {
	// Initialization of 2 Dimensional Array
	matrisk := [][]int{{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9}}
	fmt.Println("Calculate the absolute difference between the following diagonal matrix : ", matrisk)
	fmt.Println("The result of the matrix difference is : ", diagonalDiff(matrisk))
}
