package main

import (
	"encoding/json"
	"fmt"
	tags "github.com/grokify/html-strip-tags-go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

/*
   this application write stories.
*/
var usedWord = make(map[string]bool)

func main() {

	var keyword string
	var story string
	n := 100
	fmt.Scanf("%s", &keyword)
	suggestions := GetSuggestions(keyword)
	usedWord[keyword] = true
	for i := 0; i < n; i++ {
		suggestion := RandomSelect(suggestions)
		story = story + fmt.Sprintf(" %s", suggestion)
		keyword = GetRandomWord(story)
		fmt.Println(suggestion)
		suggestions = GetSuggestions(keyword)
		usedWord[keyword] = true
	}
	fmt.Println(story)
	return
}
func GetRandomWord(in string) string {
	storyslice := strings.Split(in, " ")
	ra := rand.NewSource(time.Now().UnixNano())
	ra1 := rand.New(ra)
	var nu int
	var randomWord string

	/*
		TODO
		   how to ensure that this loop does not continually go on
		   in case the limit of iterations is increased and words
		   repeat
	*/
	for randomWord == "" || usedWord[randomWord] == true {
		nu = ra1.Intn(len(storyslice))
		randomWord = storyslice[nu]
	}

	return randomWord
}

func RandomSelect(stories []string) string {
	/*
		this function accepts the slice of strings
		selects one at randome and returns it
	*/
	ra := rand.NewSource(time.Now().UnixNano())
	ra1 := rand.New(ra)
	nu := ra1.Intn(len(stories))
	return stories[nu]
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
	if len(storySuggestion) == 0 {
		fmt.Println("err:" + keyword)
	}
	return storySuggestion
}
