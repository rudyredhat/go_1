//defers the execution and the part of the code, until everything other than that is finished
//it can be used in the database connection and can be used to complete all parts and then close
// all the last connection
package main

import (
	"fmt"
)

func main() {
	// if defer keyword is added , first the open part gets print and at last the close part
	// defer keyword only works within the main function
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	myFunc()

	//LIFO demostrate
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Statement 5 Undefered")

}

func myFunc() {
	//this will print the both then it will print all the left deferred Statement of main
	defer fmt.Println("Statement rudra")
	fmt.Println("Statement rudra Undefered")
}
