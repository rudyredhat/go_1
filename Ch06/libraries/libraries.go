//group functions together in separate source code files and canbe used back as libraries
//returning multiple values from the function in the form of val & error
//libraries and stringutil folder are working together here
// ***check in the google doc for the topic = setting up go workspace for morw info

package main

import (
	"fmt"
	"stringutil" //add the import for stringutil
)

func main() {
	n1, l1 := stringutil.FullName("Rudra", "Pratap") //add the name of prefix over here
	fmt.Printf("fullname: %v, length: %v\n", n1, l1)
	n2, l3 := stringutil.FullNameNakedreturn("Rudra", "Singh")
	fmt.Printf("fullname: %v, length: %v\n", n2, l3)
}
