package main

import "fmt"

func main() {
	var p *int //this pointer can point to any int variable

	//fmt.Println("The value of p:", *p) //it will give error and can be handled using nil
	//panic: runtime error: invalid memory address or nil pointer dereference
	//[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x403569]
	if p != nil {
		fmt.Println("The value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	//point the pointer to the value of the address of the var name
	//& means connect the variable to this varibale
	p = &v
	if p != nil {
		fmt.Println("The value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1 //decalration of pointer explicitly and connect the pointer to the value
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	//but here we are changing the underlying value of the pointer1
	fmt.Println("Value 1:", value1)

}
