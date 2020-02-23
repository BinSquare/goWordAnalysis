package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExcludeStopWords(t *testing.T) {
	cases := []struct {
		input          []string
		stopWordsArray []string
		expectedOutput []string
	}{
		{
			[]string{"hello", "world"},
			[]string{"hello"},
			[]string{"world"},
		},
		{
			[]string{"hello", "world", "hello", "world"},
			[]string{"world"},
			[]string{"hello", "hello"},
		},
		{
			[]string{"then", "the", "world", "ended"},
			[]string{"the", "and", "then"},
			[]string{"world", "ended"},
		},
	}

	for index, c := range cases {
		var output []string = ExcludeStopWords(c.input, c.stopWordsArray)
		assert.Equal(t, c.expectedOutput, output, fmt.Sprintf("case %d: exclude stop words", index))
	}
}

func TestDuplicatedWord(t *testing.T) {

	cases := []struct {
		wordsList      []string
		word           string
		expectedOutput bool
	}{
		{
			[]string{"hello", "world"},
			"world",
			true,
		},
		{
			[]string{"hello", "world"},
			"again",
			false,
		},
	}

	for _, test := range cases {
		var output bool = ContainsWord(test.wordsList, test.word)
		assert.Equal(t, test.expectedOutput, output)
	}
}
