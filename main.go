package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

type stemArray struct {
	stem string
	word string
}

type wordCount struct {
	Word       string
	Occurances int
}

var stopWordsList string = "./assets/stop-words/stop_words.txt"
var lemmatizationPairs string = "./assets/lemmatization-lists/lemmatization-en.txt"

//open text file for lemmatization pairs, read & save to stemArray.
func parseStemPair(fileName string) []stemArray {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	var stemPairsArray []stemArray

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wordsArray := strings.Fields(scanner.Text())

		stem := wordsArray[0]
		lemmatizedWord := wordsArray[1]

		stemPairsArray = append(stemPairsArray,
			stemArray{stem, lemmatizedWord})

	}
	return stemPairsArray
}

//extract word stem, note - room for optimization.
func stemExtract(word string, stemPairs []stemArray) string {
	for _, pair := range stemPairs {
		if pair.word == word {
			return pair.stem
		}
	}
	return word
}

//take in a string array and return array of wordCount
func wordCounter(words []string, uniqueList []string) []wordCount {
	var tempWordCount []wordCount
	for _, uniqueWord := range uniqueList {
		counter := 0
		for _, word := range words {
			if word == uniqueWord {
				counter = counter + 1
			}
		}
		tempWordCount = append(tempWordCount, wordCount{uniqueWord, counter})
	}
	return tempWordCount
}

//open, read text file and return string of words in file
func parseText(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error:", err)
	}
	return strings.TrimRight(string(data), "\r\n")
}

//open, read text file and return array of words in file
func parseWords(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

//take in a string array and string, return bool if list contains word already.
func containsWord(wordsList []string, word string) bool {
	for _, uniqueWord := range wordsList {
		if uniqueWord == word {
			return true
		}
	}
	return false
}

//returns array with removed english stop words based on list of stopwords.
func excludeStopWords(words []string, stopWords []string) []string {
	var nonStopWords []string
	for _, word := range words {
		if !containsWord(stopWords, word) {
			nonStopWords = append(nonStopWords, word)
		}
	}
	return nonStopWords
}

//returns 25 sorted most frequent words.
func sortedWords(wordCountArray []wordCount) []wordCount {
	var tempWordCountArray []wordCount = wordCountArray
	sort.SliceStable(tempWordCountArray, func(i, j int) bool {
		return wordCountArray[i].Occurances > wordCountArray[j].Occurances
	})

	var tempSortedArray []wordCount
	for index, item := range tempWordCountArray {
		if index < 25 {
			tempSortedArray = append(tempSortedArray, item)
		}
	}
	fmt.Println(tempSortedArray)
	return tempSortedArray
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	uploadPage := template.Must(template.ParseFiles("./templates/index.html"))
	uploadPage.Execute(w, nil)
	return
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	tmpl := template.Must(template.ParseFiles("./templates/upload.html"))

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("textFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		data := wordCount{
			Word:       "ERROR, file failed to be read!",
			Occurances: 0,
		}
		tmpl.Execute(w, data)
		return
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("./test.txt", fileBytes, 0644)

	data := wordCount{
		Word: "word", Occurances: 5,
	}

	tmpl.Execute(w, data)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./templates/history.html"))
	tmpl.Execute(w, nil)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", rootHandler)
	route.HandleFunc("/upload", uploadFile)
	route.HandleFunc("/history", historyHandler)

	log.Fatal(http.ListenAndServe(":8080", route))
}
