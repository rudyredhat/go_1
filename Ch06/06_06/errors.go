//error in go defined as a instance of an interface defined as a single method named err
//which returns back string which is an error message

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.txt")
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}
	//we can override the error message
	myError := errors.New("My error string")
	fmt.Println(myError)

	//handle error instance
	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true}
	//items associated in the map with a specific key , we can use the format below
	attended, ok := attendance["Mike"] //attended will have the value from the key and ok is true, if present
	if ok {
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info from Mike")
	}
}
