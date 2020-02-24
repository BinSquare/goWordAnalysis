package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveHistory(t *testing.T) {
	testCase := []PastAnalysis{
		{
			Original: "hello world",
			Analysis: []WordCount{
				{
					Word:       "hello",
					Occurances: 1,
				},
			},
		},
	}

	createdFile := SaveHistory(testCase)
	assert.Equal(t, true, createdFile)
}

func TestReadHistory(t *testing.T) {

	expectedOutput := "hello world"
	createdFile := ReadHistory("../assets/tests/sample_history.json")
	fmt.Println(createdFile)
	assert.Equal(t, expectedOutput, createdFile[0].Original)
}
