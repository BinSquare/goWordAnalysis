package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type stemArray []struct {
	stem string
	word string
}

type wordCount []struct {
	word       string
	occurances int
}

var stopWordsList string = "./assets/stop-words/stop_words.txt"

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

func parseWords(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var stopWords []string
	for scanner.Scan() {
		stopWords = append(stopWords, scanner.Text())
	}
	return stopWords
}

//TODO take in a string array and string, return bool if list contains word already.
func duplicatedWord(wordsList []string, word string) bool {
	for _, uniqueWord := range wordsList {
		if uniqueWord == word {
			return true
		}
	}
	return false
}

//TODO exclude english stop words
func excludeStopWords(words []string, stopWords []string) []string {
	var nonStopWords []string
	for _, word := range words {
		for _, stopWord := range stopWords {
			if stopWord != word && !duplicatedWord(nonStopWords, word) {
				nonStopWords = append(nonStopWords, word)
			}
		}
	}
	return nonStopWords
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	uploadPage := template.Must(template.ParseFiles("./templates/index.html"))
	uploadPage.Execute(w, nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", rootHandler)
	route.HandleFunc("/upload", uploadHandler)

	log.Fatal(http.ListenAndServe(":8080", route))
}
