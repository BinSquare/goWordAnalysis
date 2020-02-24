package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

//StemArray returns stem and lemmatized word pair
type StemArray struct {
	stem string
	word string
}

//ParseStemPair open text file for lemmatization pairs, read & save to stemArray.
func ParseStemPair(fileName string) []StemArray {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	var stemPairsArray []StemArray

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wordsArray := strings.Fields(scanner.Text())

		stem := wordsArray[0]
		lemmatizedWord := wordsArray[1]

		stemPairsArray = append(stemPairsArray,
			StemArray{stem, lemmatizedWord})

	}
	return stemPairsArray
}

//StemExtract extract word stem, note - room for optimization.
func StemExtract(word string, stemPairs []StemArray) string {
	for _, pair := range stemPairs {
		if pair.word == word {
			return pair.stem
		}
	}
	return word
}

//Stemify iterates words array to return a array with only word stems.
func Stemify(wordsList []string, stemPairsList []StemArray) []string {
	var stemmedList []string
	for _, word := range wordsList {
		stemmedList = append(stemmedList, StemExtract(word, stemPairsList))
	}
	return stemmedList
}
