package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseText(t *testing.T) {
	var output string = parseText("../assets/tests/sample_input_2.txt")
	assert.Equal(t, "hello world!", output)
}

func TestParseStemPairs(t *testing.T) {

	stemPair := parseStemPair("../assets/tests/sample_lemmatization_pairs_1.txt")

	stemtestArray := []stemArray{
		{"play", "playing"},
		{"stay", "staying"},
	}
	assert.Equal(t, stemtestArray, stemPair, fmt.Sprintf("Extracts stem pairs from file."))
}
