package main

import "fmt"

func main() {
	var x float64 = 42
	var result string

	if x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "= to 0"
	} else {
		result = "greater to zero"
	}
	fmt.Println("Result:", result)
  //if the value of x is defined in the if block
  //it cannot be used outside the loop
  // if x := 42; x < 0 //so we cannot use this outside if block

}
