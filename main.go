package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	utils "github.com/BinSquare/goWordAnalysis/vendor"
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

	os.MkdirAll("./uploads/", os.ModePerm)
	fileName := fmt.Sprintf("./uploads/%d.txt", time.Now().UnixNano())
	ioutil.WriteFile(fileName, fileBytes, 0644)

	fileContent := utils.ParseText(fileName)

	sortedList := wordAnalysis(fileName, lemmatizationPairs, filter, stopWordsList)
	_ = os.Remove(fileName)

	//spaghetti code from this point forward, out of time.
	histories := utils.ReadHistory("./history.json")

	history := utils.PastAnalysis{
		Original: fileContent,
		Analysis: sortedList,
	}

	if len(histories) > 10 {
		histories = histories[1:]
	}

	histories = append(histories, history)

	utils.SaveHistory(histories)
	data := PageData{
		FileContent: fileContent,
		FileWords:   sortedList,
		Filter:      filter,
	}

	tmpl.Execute(w, data)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {

	histories := utils.ReadHistory("./history.json")
	fmt.Println(histories)
	type HistoryData struct {
		Content []utils.PastAnalysis
	}

	data := HistoryData{
		Content: histories,
	}

	fmt.Println(data)
	tmpl := template.Must(template.ParseFiles("./templates/history.html"))
	tmpl.Execute(w, data)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", rootHandler)
	route.HandleFunc("/upload", uploadFile)
	route.HandleFunc("/history", historyHandler)

	type Analysis struct {
		Original string
		Analysis []utils.WordCount
	}

	log.Fatal(http.ListenAndServe(":8080", route))
}
