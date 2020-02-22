package utils

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

	stemPairs := ParseStemPair("../assets/lemmatization-lists/lemmatization-en.txt")

	for index, c := range cases {
		var output string = StemExtract(c.input, stemPairs)
		assert.Equal(t, c.expectedOutput, output, fmt.Sprintf("Case%d: %s returned stem of %s", index, c.input, output))
	}
}
