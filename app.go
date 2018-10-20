package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

/*
   this application write stories.
*/
func main() {
	/*
		get the first letter
		make a google reques
		get the next letter
		if number of words is greater than
		word limit, exit the program
	*/
	var keyword string
	clientUrl := os.Getenv("CLI")
	fmt.Scanf("%s", &keyword)
	r, err := http.Get(clientUrl + keyword)
	if err != nil {
		panic(err)
	}
	bytestring, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var wordMap []interface{}
	err = json.Unmarshal([]byte(bytestring), &wordMap)
	if err != nil {
		panic(err)
	}
	for _, val := range wordMap {
		if rt := reflect.TypeOf(val); rt.Kind() == reflect.Slice {
			for _, val1 := range val.([]interface{}) {
				fmt.Println(val1)
			}

		}
	}
	return
}
