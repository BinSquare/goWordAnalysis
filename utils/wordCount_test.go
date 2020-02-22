package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		input          []string
		uniqueList     []string
		expectedOutput []wordCount
	}{
		{
			[]string{
				"hello", "world",
			},
			[]string{
				"hello", "world",
			},
			[]wordCount{
				{"hello", 1},
				{"world", 1},
			},
		},
		{
			[]string{
				"hello", "world", "hello", "hello", "world",
			},
			[]string{
				"hello", "world",
			},
			[]wordCount{
				{"hello", 3},
				{"world", 2},
			},
		},
	}

	for index, c := range cases {
		var output []wordCount = wordCounter(c.input, c.uniqueList)
		assert.Equal(t, c.expectedOutput, output, fmt.Sprintf("case %d: input %s", index, c.input))
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
		var output bool = containsWord(test.wordsList, test.word)
		assert.Equal(t, test.expectedOutput, output)
	}
}

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
		var output []string = excludeStopWords(c.input, c.stopWordsArray)
		assert.Equal(t, c.expectedOutput, output, fmt.Sprintf("case %d: exclude stop words", index))
	}
}

func TestSortedWords(t *testing.T) {
	cases := []struct {
		input          []wordCount
		expectedOutput []wordCount
	}{
		{
			[]wordCount{
				{"you", 49},
				{"you", 49},
				{"you", 49},
				{"you", 49},
				{"you", 49},
				{"ends", 4},
				{"ends", 4},
				{"ends", 4},
				{"ends", 4},
				{"ends", 4},
				{"world", 1},
				{"world", 1},
				{"world", 1},
				{"world", 1},
				{"world", 1},
				{"what", 0},
				{"the", 850},
				{"the", 850},
				{"the", 850},
				{"the", 850},
				{"the", 850},
				{"with", 150},
				{"with", 150},
				{"with", 150},
				{"with", 150},
				{"with", 150},
			},
			[]wordCount{
				{"the", 850},
				{"the", 850},
				{"the", 850},
				{"the", 850},
				{"the", 850},
				{"with", 150},
				{"with", 150},
				{"with", 150},
				{"with", 150},
				{"with", 150},
				{"you", 49},
				{"you", 49},
				{"you", 49},
				{"you", 49},
				{"you", 49},
				{"ends", 4},
				{"ends", 4},
				{"ends", 4},
				{"ends", 4},
				{"ends", 4},
				{"world", 1},
				{"world", 1},
				{"world", 1},
				{"world", 1},
				{"world", 1},
			}},
	}
	for _, c := range cases {
		output := sortedWords(c.input)
		assert.Equal(t, c.expectedOutput, output)
	}

}
