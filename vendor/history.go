package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//PastAnalysis is a struct containing old runs of original & wordCount analysis.
type PastAnalysis struct {
	Original string
	Analysis []WordCount
}

// FileExists returns if file exists and is not a directory.
func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// SaveHistory creates file and return bool on success/fail of file creation.
func SaveHistory(data []PastAnalysis) bool {
	file, _ := json.MarshalIndent(data, "", " ")

	err := ioutil.WriteFile("history.json", file, 0644)

	if err != nil {
		return false
	}

	return true
}

// ReadHistory reads file and returns PastAnalytics struct
func ReadHistory(fileName string) []PastAnalysis {
	file, _ := ioutil.ReadFile(fileName)

	data := []PastAnalysis{}

	_ = json.Unmarshal([]byte(file), &data)
	return data
}
