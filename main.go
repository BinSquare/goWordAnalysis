package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	utils "goWordAnalysis/vendor"

	"github.com/gorilla/mux"
)

var stopWordsList string = "./assets/stop-words/stop_words.txt"
var lemmatizationPairs string = "./assets/lemmatization-lists/lemmatization-en.txt"

func wordAnalysis(fileName string, lemmaFile string, filterStopWords bool, stopWordslist string) []utils.WordCount {
	ogWords := utils.ParseWords(fileName)
	stopWords := utils.ParseWords(stopWordsList)

	stemPairsList := utils.ParseStemPair(lemmaFile)

	stemmedList := utils.Stemify(ogWords, stemPairsList)
	uniqueList := utils.Uniqify(stemmedList)

	if filterStopWords {
		filteredList := utils.ExcludeStopWords(uniqueList, stopWords)
		wordCountList := utils.WordCounter(stemmedList, filteredList)
		sortedList := utils.SortedWords(wordCountList)
		return sortedList
	}
	wordCountList := utils.WordCounter(stemmedList, uniqueList)
	sortedList := utils.SortedWords(wordCountList)
	return sortedList
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	uploadPage := template.Must(template.ParseFiles("./templates/index.html"))
	uploadPage.Execute(w, nil)
	return
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	tmpl := template.Must(template.ParseFiles("./templates/upload.html"))

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("textFile")
	checkbox := r.Form["stopWordsFilter"]
	filter := false
	if checkbox != nil {
		filter = true
	}

	type PageData struct {
		FileContent string
		FileWords   []utils.WordCount
		Filter      bool
	}

	fmt.Println(checkbox)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		data := PageData{
			FileContent: "ERROR, file failed to be read!",
			Filter:      filter,
		}
		tmpl.Execute(w, data)
		return
	}

	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fileName := fmt.Sprintf("./uploads/%d.txt", time.Now().UnixNano())
	ioutil.WriteFile(fileName, fileBytes, 0644)

	fileContent := utils.ParseText(fileName)

	sortedList := wordAnalysis(fileName, lemmatizationPairs, filter, stopWordsList)

	data := PageData{
		FileContent: fileContent,
		FileWords:   sortedList,
		Filter:      filter,
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
