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
		expectedOutput []WordCount
	}{
		{
			[]string{
				"hello", "world",
			},
			[]string{
				"hello", "world",
			},
			[]WordCount{
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
			[]WordCount{
				{"hello", 3},
				{"world", 2},
			},
		},
	}

	for index, c := range cases {
		var output []WordCount = WordCounter(c.input, c.uniqueList)
		assert.Equal(t, c.expectedOutput, output, fmt.Sprintf("case %d: input %s", index, c.input))
	}
}

func TestSortedWords(t *testing.T) {
	cases := []struct {
		input          []WordCount
		expectedOutput []WordCount
	}{
		{
			[]WordCount{
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
			[]WordCount{
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
		output := SortedWords(c.input)
		assert.Equal(t, c.expectedOutput, output)
	}

}
