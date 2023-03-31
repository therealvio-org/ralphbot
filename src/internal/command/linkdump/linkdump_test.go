package linkdump

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinks(t *testing.T) {

	foo := "foo - <foo.com>"
	bar := "bar - <bar.com>"

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "when one link is provided as input, the string contains one link",
			input: `{
				"Links": [
					{
					  "Name": "foo",
					  "URL": "foo.com"
					}
				]
			}`,
			expected: foo,
		},
		{
			name: "when 2 links are provided as input, the string contains two links, seperated by newline",
			input: `{
				"Links": [
					{
					  "Name": "foo",
					  "URL": "foo.com"
					},
					{
					  "Name": "bar",
					  "URL": "bar.com"
					}
				]
			}`,
			expected: fmt.Sprintf("%s\n%s", foo, bar),
		},
	}

	for _, test := range cases {
		result, err := getLinks([]byte(test.input))
		assert.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}
