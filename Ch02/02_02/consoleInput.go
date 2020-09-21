package main

import (
	"bufio" //used for multiple processing
	"fmt"
	"os"
	"strconv" //strconv will not work, and will not parse string
	//strconv.ParseFloat: parsing "dsa ads\n": invalid syntax
	//so we need use another package
	"strings" //use of TrimSpace and remove the whitespace
)

func main() {
	// var_keyword  var_name data_type --> explicit typed declaration
	var s string
	//Scanln can be passed as pass by reference
	//it parse the string the first value is used not the set of values
	fmt.Scanln(&s)
	fmt.Println(s)

	//reader obj can take variety of inputs
	//below reader obj can take multiple value inputs from Stdin
	reader := bufio.NewReader(os.Stdin) // --> implicit typed declaration with :=
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	//if we want to use numeric var
	//if the type is already declared we can not re declare it
	fmt.Print("Enter a number: ")
	//assign the values to pre declare type
	str, _ = reader.ReadString('\n')
	//string now must be converted to a number, create a new var
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}

}
