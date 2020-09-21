// every file has a package declaration with a keyword package
// start up file always has a package of main , its an identifier not a string
package main

// go-lib are organised as packages, but we have to explicitly import everything
import (
	"fmt"
	"strings"
)

// always have a func named as main
func main() {
	fmt.Println("Hello from Go")
	fmt.Println(strings.ToUpper("hello really lol!"))
}

// open terminal and run -- go run helloworld.go
