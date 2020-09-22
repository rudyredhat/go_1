//function can be created
package main

import "fmt"

func main() {
	doSomething()
	sum := addValues(23, 54)
	fmt.Println("Sum:", sum)
	sum = addAddValues(12, 54, 79)
	fmt.Println("New Sum:", sum)
}

func doSomething() {
	fmt.Println("Doing Something!")
}

func addValues(value1 int, value2 int) int { //take the 2 values and at last one 1 int return type
	return value1 + value2
}

//multiple values can be processed in the single function
func addAddValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
