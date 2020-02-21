package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStemExtract(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput string
	}{
		{"", ""},
		{"talks", "talk"},
		{"talking", "talk"},
		{"talked", "talk"},
		{"plays", "play"},
		{"playing", "play"},
		{"played", "play"},
		{"passes", "pass"},
		{"passing", "pass"},
		{"passed", "pass"},
		{"copies", "copy"},
		{"copied", "copy"},
		{"copying", "copy"},
	}

	for index, c := range cases {
		var output string = stemExtract(c.input)
		assert.Equal(t, output, c.expectedOutput, fmt.Sprintf("Case%d: %s returned stem of %s", index, c.input, output))
	}
}

func TestParseStemPairs(t *testing.T) {

	stemPair := parseStemPair("test.txt")
	assert.Equal(t, stemPair, stemArray{{"test", "testing"}}, fmt.Sprintf("Extracts stem pairs from file."))
}

func TestWordCount(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput wordCount
	}{
		{
			"hello world",
			wordCount{
				{"hello", 1},
				{"world", 1},
			},
		},
	}

	for index, c := range cases {
		var output wordCount = wordCounter("hello world")
		assert.Equal(t, output, c.expectedOutput, fmt.Sprintf("case %d: input %s", index, c.input))
	}
}

func TestParseText(t *testing.T) {
	var output string = parseText("assets/tests/sample_input_2.txt")
	assert.Equal(t, "hello world!", output)
}

// func TestUniqueSet(t *testing.T) {
// 	var output []string = uniqueSet("what")
// 	assert.Equal(t, []string{"hello", "world"}, output)
// }

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
		var output bool = duplicatedWord(test.wordsList, test.word)
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
			[]string{"hello", "world"},
			[]string{"world"},
			[]string{"hello"},
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
