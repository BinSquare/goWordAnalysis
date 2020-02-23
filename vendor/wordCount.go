package utils

import (
	"fmt"
	"sort"
)

//WordCount structure contains both the word and occurances of the word
type WordCount struct {
	Word       string
	Occurances int
}

//WordCounter take in a string array and return array of wordCount
func WordCounter(words []string, uniqueList []string) []WordCount {
	var tempWordCount []WordCount
	for _, uniqueWord := range uniqueList {
		counter := 0
		for _, word := range words {
			if word == uniqueWord {
				counter = counter + 1
			}
		}
		tempWordCount = append(tempWordCount, WordCount{uniqueWord, counter})
	}
	return tempWordCount
}

//SortedWords returns 25 sorted most frequent words.
func SortedWords(wordCountArray []WordCount) []WordCount {
	var tempWordCountArray []WordCount = wordCountArray
	sort.SliceStable(tempWordCountArray, func(i, j int) bool {
		return wordCountArray[i].Occurances > wordCountArray[j].Occurances
	})

	var tempSortedArray []WordCount
	for index, item := range tempWordCountArray {
		if index < 25 {
			tempSortedArray = append(tempSortedArray, item)
		}
	}
	fmt.Println(tempSortedArray)
	return tempSortedArray
}
