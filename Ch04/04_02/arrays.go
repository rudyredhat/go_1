package main

import (
	"fmt"
)

//better to use slices rather than array
func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Yellow"
	colors[2] = "Green"
	fmt.Println(colors)
	//[Red Yellow Green]

	var numbers = [5]int{5, 3, 2, 1, 4}
	fmt.Println(numbers, len(numbers))
	//[5 3 2 1 4]

}
