package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("launching at: %s\n", t)
	//launching at: 2009-11-10 23:00:00 +0000 UTC

	now := time.Now()
	fmt.Printf("System: %s\n", now)
	//System: 2020-09-21 23:18:26.750741858 +0530 IST m=+0.000341437

	fmt.Println("Month:", t.Month())
	fmt.Println("Day:", t.Day())
	fmt.Println("weekday:", t.Weekday())

	//to check the next Date , better approach to have constant available online using .Format() *
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v %v %v %v\n",
		tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

}
