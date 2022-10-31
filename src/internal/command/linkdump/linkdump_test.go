package linkdump

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinks(t *testing.T) {

	foo := "foo - <foo.com>"
	bar := "bar - <bar.com>"

	linkJson := `{
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
	}`

	expected := fmt.Sprintf("%s\n%s", foo, bar)
	links, err := getLinks([]byte(linkJson))

	assert.NoError(t, err)
	assert.Equal(t, expected, links)
}
