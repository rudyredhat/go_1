package main

import "fmt"

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true
	// concatenate the string in the single char
	// 	fmt.Println(str1, str2, str3)
	// can be allowed to return more than 1 values simulatenously
	stringLength, err := fmt.Println(str1, str2, str3)
	// if the err values was not checked using if statement there will be an error
	// to flow the error variable we can use (_) in the starting
	//  stringLength, _ := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String Length: ", stringLength)
	}
	// Printf used for placeholders called verb , %v
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	// wrapping a int to float
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))
	// get the types of each var type
	fmt.Printf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	// changing to outputing to the string, changing the function to Sprintf which means return the string
	myString := fmt.Sprintf("Data Types as var : %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)

}
