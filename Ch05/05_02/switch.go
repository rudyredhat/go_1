package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //seed the time to get the miliseconds from local computer
	dow := rand.Intn(6) + 1      //provides the random integer value from 0 to ceiling value , i.e 6 and then we are adding 1 to it
	fmt.Println("Day", dow)

	result := ""
	switch dow {
	case 1:
		result = "Sunday"
	case 7:
		result = "Saturday"
	default:
		result = "Weekday"
	}
	fmt.Println("Day", dow, "Result", result)

	x := 42 //alternative of if statement
	switch {
	case x < 0:
		result = "x<0"
		//fallthrough  ****imp***
		//it same as break keyword
		//and it will reevaluate the value of x to 0 and print the next case
		//so both above and below part is checked
	case x == 0:
		result = "x=0"
	default:
		result = "x>0"
	}
	fmt.Println(result)

}
