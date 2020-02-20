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

func TestUniqueSet(t *testing.T) {
	var output []string = uniqueSet("what")
	assert.Equal(t, []string{"hello", "world"}, output)
}
