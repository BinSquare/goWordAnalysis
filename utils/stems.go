package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type stemArray struct {
	stem string
	word string
}

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
