//getting back the http response and the http background code (html, json ) etc
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://bugnoobie.wordpress.com/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//if error is not occured we can check the response type
	fmt.Printf("Response type: %T\n", resp)

	//no the body part
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//but still it is byte arrary need to convert it to string
	content := string(bytes)
	fmt.Print(content)

}
