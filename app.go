package main

import (
	"encoding/json"
	"fmt"
	tags "github.com/grokify/html-strip-tags-go"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

/*
   this application write stories.
*/
func main() {

	var keyword string
	fmt.Scanf("%s", &keyword)
	suggestions := GetSuggestions(keyword)
	fmt.Println(suggestions)
	return
}
func RandomSelect(stories []string) string {

}
func GetSuggestions(keyword string) []string {
	/*
		get the first letter
		make a google reques
		get the next letter
		if number of words is greater than
		word limit, exit the program
	*/
	clientUrl := os.Getenv("CLI")
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
	var storySuggestion []string
	for _, val := range wordMap {
		if rt := reflect.TypeOf(val); rt.Kind() == reflect.Slice {
			for _, val1 := range val.([]interface{}) {
				suggestion := val1.([]interface{})[0]
				clearSuggestion := tags.StripTags(suggestion.(string))
				storySuggestion = append(storySuggestion, clearSuggestion)
			}

		}
	}
	return storySuggestion
}
