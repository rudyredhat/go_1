package main

import (
	"fmt"
	"sort"
)

//slice is an abstraction layer on top of array
//when declared slice = runtime declared underlying array = allocates the memory
// then returns back the requested slice
func main() {
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	//since slices are resizable we can add and remove item from slices
	colors = append(colors, "purple")
	fmt.Println(colors)
	//remove item from slices
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	//we can have the default as the length of colors
	colors = append(colors[1:])
	fmt.Println(colors)
	//remove the last ele
	colors = append(colors[0 : len(colors)-1])
	fmt.Println(colors)

	//initialise the slice with make function ***imp***
	numbers := make([]int, 5, 5) //type, size, capacity
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 4
	numbers[3] = 5
	numbers[4] = 3
	fmt.Println(numbers)

	//appending a new item will incerase the capacity at the run runtime
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	//sort a slice = ascending , add the sort package first
	sort.Ints(numbers)
	fmt.Println(numbers)

}
