package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

//here the cutom type dog is the interface for the Animal
//below we have 3 types all implementing the same interface

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meeow!"
}

type Cow struct {
}

//here the cutom tyoe dog is the interface for the Animal
func (c Cow) Speak() string {
	return "Moo!"
}

func main() {
	//we need to caste the object when it is declared
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
