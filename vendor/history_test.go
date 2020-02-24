package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveHistory(t *testing.T) {
	cases := []PastAnalysis{
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

	createdFile := SaveHistory(cases)
	assert.Equal(t, true, createdFile)
}

func TestReadHistory(t *testing.T) {
	createdFile := ReadHistory("../assets/tests/sample_history.json")
	fmt.Println(createdFile)
	assert.Equal(t, true, createdFile)
}
