package main

import (
	"fmt"
	"net/http"
)

type stemArray []struct {
	stem string
	word string
}

//TODO open text file for lemmatization pairs, read & save to stemArray.
func parseText(fileName string) stemArray {
	testStemArray := stemArray{
		{"one", "ones"},
	}
	return testStemArray
}

//TODO extract word stem
func stemExtract(word string) string {
	return "test"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
