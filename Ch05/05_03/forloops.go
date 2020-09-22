package main

import (
	"fmt"
)

func main() {
	sum := 1
	fmt.Println("Sum:", sum)
	colors := []string{"red", "green", "blue"}
	fmt.Println(colors)

	sum = 0 //reset the value to 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Sum:", sum)
	}
	fmt.Println("Final Sum:", sum)

	//loop through the slice of values
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	//iterating through the range keyword
	//set the value of i to the current index and loop through the values of i
	for i := range colors {
		fmt.Println(colors[i])
	}

	//there is no such while loop, but can use for loop in the same way
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}

	//labels can created to jump to the code
endofprogram:
	fmt.Println("end of program")
}
