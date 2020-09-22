//returning multiple values from the function in the form of val & error
package main

import "fmt"

func main() {
	n1, l1 := FullName("Rudra", "Pratap")
	fmt.Printf("fullname: %v, length: %v\n", n1, l1)
	n2, l3 := FullNameNakedreturn("Rudra", "Singh")
	fmt.Printf("fullname: %v, length: %v\n", n2, l3)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

//for naked retrun, without taking care of the return order

func FullNameNakedreturn(f, l string) (full string, length int) {
	full = f + "" + l
	length = len(full)
	return
}
