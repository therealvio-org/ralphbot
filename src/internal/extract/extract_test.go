package extract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectRandomString(t *testing.T) {
	cases := []struct {
		name  string
		input []string
	}{
		{
			name: "given 3 elements in a slice, SelectRandomString will randomly pick an element",
			input: []string{
				"option 1",
				"option 2",
				"option 3",
			},
		},
	}

	for _, test := range cases {
		result := RandomString(test.input)

		// testing for probability i.e. does our result show up at least once? This is more of a "transparency" test
		var resultTally []string
		for i := 0; i < 1000; i++ {
			resultTally = append(resultTally, RandomString(test.input))
		}

		assert.Contains(t, resultTally, result)
	}
}
