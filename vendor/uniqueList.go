package utils

//ContainsWord take in a string array and string, return bool if list contains word already.
func ContainsWord(wordsList []string, word string) bool {
	for _, uniqueWord := range wordsList {
		if uniqueWord == word {
			return true
		}
	}
	return false
}

//Uniqify removes duplicate words from list and returns unique list.
func Uniqify(wordsList []string) []string {
	var uniqueList []string
	for _, word := range wordsList {
		if !ContainsWord(uniqueList, word) {
			uniqueList = append(uniqueList, word)
		}
	}
	return uniqueList
}
