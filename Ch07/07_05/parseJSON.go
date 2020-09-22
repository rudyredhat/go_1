//getting back the http response and the http background code (html, json ) etc
//pasrse the text and turn down to structured data
//first go the url and identify the fields we want to retrieve back

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big" //format the final data in proper manner
	"net/http"
	"strings"
)

type Tour struct { //take this value from json
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	tours := toursFromJson(content)
	fmt.Println(tours)
	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {

	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	//i do not know how many tours how we will get

	decoder := json.NewDecoder(strings.NewReader(content))
	//reader will do the content and will decoder will decode it
	_, err := decoder.Token()
	//will remove the first token from the decoder package
	checkError(err)

	var tour Tour
	//loop through the content of the json package
	for decoder.More() { //if there is another data enity . loop through one by one
		err := decoder.Decode(&tour)
		//Decode method is doing the main part, parse the json content and will grab the fields defined
		checkError(err)
		//if no error
		tours = append(tours, tour)
		//append the tours object
	}
	return tours
}
