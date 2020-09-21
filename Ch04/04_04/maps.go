//a map in go is collection of key value pair
//it is just a hash table and datas can be retrieved from the key

package main

import (
	"fmt"
	"sort"
)

func main() {
	//empty map using make function
	states := make(map[string]string) //key and values both are string
	fmt.Println(states)

	//separated by space char , similar to dict in python
	states["B"] = "patna"
	states["WB"] = "kolkata"
	states["RJ"] = "jaipur"
	fmt.Println(states)

	//retrive an item from map
	patna := states["B"]
	fmt.Println(patna)

	//delete an item from the map
	delete(states, "WB")
	fmt.Println(states)

	//iterate through the map
	states["K"] = "bangalore"
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
		//maps iteration order is not same always
	} //range_keyword map_name
	//to make it orderd, extract the keys from the map as a slice of strings
	//1) created a slice and populated keys from the map
	keys := make([]string, len(states))
	i := 0
	//using range keyword having one value will return only key
	for k := range states {
		keys[i] = k
		i++
	}
	//sorted that slice
	sort.Strings(keys)
	//report the results through iterating through the slice
	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
