package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type stemArray []struct {
	stem string
	word string
}

type wordCount []struct {
	word       string
	occurances int
}

//TODO open text file for lemmatization pairs, read & save to stemArray.
func parseStemPair(fileName string) stemArray {
	testStemArray := stemArray{
		{"one", "ones"},
	}
	return testStemArray
}

//TODO extract word stem
func stemExtract(word string) string {
	return "test"
}

//TODO take in a string and return array of wordCount
func wordCounter(words string) wordCount {
	return wordCount{{"test", 1}}
}

//TODO open, read text file and return string of words in file
func parseText(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error:", err)
	}
	return string(data)
}

//TODO take in a string and return a unique array of words.
func uniqueSet(list string) []string {
	return []string{"hello", "world"}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
