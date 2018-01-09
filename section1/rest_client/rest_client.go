// Package restclient is a simple package
// that fetches json page from the url
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "http://quotes.rest/qod.json"

func main() {
	res, err := http.Get(apiURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
