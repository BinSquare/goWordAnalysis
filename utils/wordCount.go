package utils

import (
	"fmt"
	"sort"
)

type wordCount struct {
	Word       string
	Occurances int
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
