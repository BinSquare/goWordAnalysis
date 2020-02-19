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

	for _, c := range cases {
		var output string = stemExtract(c.input)
		assert.Equal(t, output, c.expectedOutput, fmt.Sprintf("The word %s returned stem of %s", c.input, output))

	}
}
