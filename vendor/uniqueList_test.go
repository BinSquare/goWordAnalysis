package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestUniqify(t *testing.T) {
	cases := []struct {
		wordsList      []string
		expectedOutput []string
	}{
		{
			[]string{"hello", "world", "world"},
			[]string{"hello", "world"},
		},
		{
			[]string{"hello", "hello", "hello"},
			[]string{"hello"},
		},
	}

	for _, test := range cases {
		var output []string = Uniqify(test.wordsList)
		assert.Equal(t, test.expectedOutput, output)
	}
}
