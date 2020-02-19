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

func TestTextParser(t *testing.T) {

	stemPair := parseText("test.txt")
	assert.Equal(t, stemPair, stemArray{{"test", "testing"}}, fmt.Sprintf("Extracts stem pairs from file."))
}
