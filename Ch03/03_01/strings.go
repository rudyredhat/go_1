package main

import (
	"fmt"
	//strings can be declared as single simple val or aaray of bytes
	"strings"
)

func main() {
	str1 := "An implicitly types string"
	fmt.Println(str1)
	//to check the actual type
	fmt.Printf("String:type = %v:%T\n", str1, str1)
	str2 := "An explicit types string"
	fmt.Println(str2)
	//to check the actual type
	fmt.Printf("String:type = %v:%T\n", str2, str2)

	//extra strings packages functions
	//uppercase
	fmt.Println(strings.ToUpper(str1))
	//change the first char to capitalise
	fmt.Println(strings.Title(str1))

	//comparion of two strings
	lvalue := "hello"
	uvalue := "HELLO"
	fmt.Println("Equal?", (lvalue == uvalue)) //it will return false
	//need to check with non-string case comparison
	fmt.Println("Equal NOn-case compare?", strings.EqualFold(lvalue, uvalue)) //return true

	//looks for a value in another value
	fmt.Println("Contain exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contain exp?", strings.Contains(str2, "exp"))

}
