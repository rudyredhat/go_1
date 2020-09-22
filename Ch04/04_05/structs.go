//struct type in go is a data structure
//similar to purpose of functions
//each struct is idependent of itself
//in this focus on fields for data
package main

import "fmt"

type Dog struct { //type struct_name struct_keyword
	Breed string //fields are defined here, if wanted to be publicly accessible have fields
	// in starting letter as capital
	Weight int
}

func main() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	//{Poodle 34}
	//we can dump the whole result in the same fields
	fmt.Printf("%+v\n", poodle)
	//each value in the structure is a field
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}
