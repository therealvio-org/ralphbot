package linkdump

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}

func TestGetLinksFromJSON(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected LinksStruct
	}{
		{
			name: "when one link is provided as input, the struct contains one unmarshalled link",
			input: []byte(`{
				"links": [
					{
					  "Name": "foo",
					  "URL": "foo.com"
					}
				]
			}`),
			expected: LinksStruct{
				Links: []LinkDetails{
					{
						Name: "foo",
						URL:  "foo.com",
					},
				},
			},
		},
		{
			name: "when one link is provided as input, the struct contains one unmarshalled link",
			input: []byte(`{
				"links": [
					{
					  "Name": "foo",
					  "URL": "foo.com"
					},
					{
					  "Name": "bar",
					  "URL": "bar.com"
					}
				]
			}`),
			expected: LinksStruct{
				Links: []LinkDetails{
					{
						Name: "foo",
						URL:  "foo.com",
					},
					{
						Name: "bar",
						URL:  "bar.com",
					},
				},
			},
		},
	}

	for _, test := range cases {
		result, err := getLinksFromJSON([]byte(test.input))
		assert.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestMakeLinkDumpMessage(t *testing.T) {
	foo := "foo - <foo.com>"
	bar := "bar - <bar.com>"

	cases := []struct {
		name   string
		input  LinksStruct
		expect string
	}{
		{
			name: "when the input LinksStruct contains one link, return one link in the string",
			input: LinksStruct{
				Links: []LinkDetails{
					{
						Name: "foo",
						URL:  "foo.com",
					},
				},
			},
			expect: fmt.Sprintf("%v", foo),
		},
		{
			name: "when the input LinksStruct contains two links, return two links in the string, seperated by newline",
			input: LinksStruct{
				Links: []LinkDetails{
					{
						Name: "foo",
						URL:  "foo.com",
					},
					{
						Name: "bar",
						URL:  "bar.com",
					},
				},
			},
			expect: fmt.Sprintf("%v\n%v", foo, bar),
		},
	}

	for _, test := range cases {
		result := makeLinkDumpMessage(test.input)
		assert.Equal(t, test.expect, result)
	}
}
