package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./fromSting.txt") //file doed not exist, but can be act
	// as a pointer to that file
	checkError(err) //if any error is found, we wont pass this stage
	defer file.Close()

	//write to the file
	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with file of %v characters", ln)

	//writes byte array instead of simple string
	bytes := []byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)
}

func checkError(err error) {
	if err != nil {
		panic(err) //show the error and halts the program
	}

}
