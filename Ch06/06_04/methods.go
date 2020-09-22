package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

//in java like each method is a member of a class
//in go method is a member of a type

func (d Dog) Speak() { //keyword_func var_name data_tye_attaching_var_to method_name
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() { //keyword_func var_name data_tye_attaching_var_to method_name
	d.Sound = fmt.Sprintf("%v! %v! %v!\n ", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}

func (d *Dog) SpeakThreeTimesCheck() { //keyword_func var_name data_tye_attaching_var_to method_name
	d.Sound = fmt.Sprintf("%v! %v! %v! ", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()

	//can change the contents of the poodle objects
	poodle.Sound = "Arf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	//but when we print this again we are receiving d as an original value
	poodle.SpeakThreeTimes()

	//now it will work as pointer, so the changes that are making in the function are affecting the original obj
	poodle.SpeakThreeTimesCheck()
	poodle.SpeakThreeTimesCheck()

}
