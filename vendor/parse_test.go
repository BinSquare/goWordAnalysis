package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseText(t *testing.T) {
	var output string = ParseText("../assets/tests/sample_input_2.txt")
	assert.Equal(t, "hello world!", output)
}

func TestParseStemPairs(t *testing.T) {

	stemPair := ParseStemPair("../assets/tests/sample_lemmatization_pairs_1.txt")

	stemtestArray := []StemArray{
		{"play", "playing"},
		{"stay", "staying"},
	}
	assert.Equal(t, stemtestArray, stemPair, fmt.Sprintf("Extracts stem pairs from file."))
}
