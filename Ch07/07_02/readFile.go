package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "./hello.txt"
	content, err := ioutil.ReadFile(fileName)
	checkError(err)

	fmt.Println("read from file:", content)
	// it is dumping the raw bytes
	// read from file: [72 101 108 108 111 32 119 111 114 108 100 10]
	// now need to take the byte array and convert it to string

	result := string(content)
	fmt.Println("read from file:", result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
