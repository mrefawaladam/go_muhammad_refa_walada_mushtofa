package main 
 
  import (
	"fmt"
  )

  func main() {
	// APPEND & COPY
	var colors = []string{"black", "white", "RED", "GREEN"}
	colors = append(colors,"blue")

	copied_color := make([]string,10)

	copy(copied_color, colors)
	fmt.Println(copied_color)
  }