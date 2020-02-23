package utils

//ExcludeStopWords returns array with removed english stop words based on list of stopwords.
func ExcludeStopWords(words []string, stopWords []string) []string {
	var nonStopWords []string
	for _, word := range words {
		if !ContainsWord(stopWords, word) {
			nonStopWords = append(nonStopWords, word)
		}
	}
	return nonStopWords
}

//ContainsWord take in a string array and string, return bool if list contains word already.
func ContainsWord(wordsList []string, word string) bool {
	for _, uniqueWord := range wordsList {
		if uniqueWord == word {
			return true
		}
	}
	return false
}
