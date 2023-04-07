package dadjoke

import (
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

func TestGetJokes(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect struct {
			result struct {
				slice        []string
				errorMessage string
			}
		}
	}{
		{
			name: "with a json array length of 3, getJokes returns a slice with length of 3",
			input: `{
						"jokes": [
							"the joke is this test suite",
							"not really",
							"unless?"
						]
					}`,
			expect: struct {
				result struct {
					slice        []string
					errorMessage string
				}
			}{
				result: struct {
					slice        []string
					errorMessage string
				}{
					slice: []string{"the joke is this test suite", "not really", "unless?"},
				},
			},
		},
		{
			name: "when a bad json input is provided, getJokes returns a slice with length of 0, and an error",
			input: `{
				"bad-json": {
					notAValidKey: true
				}
			}`,
			expect: struct {
				result struct {
					slice        []string
					errorMessage string
				}
			}{
				result: struct {
					slice        []string
					errorMessage string
				}{
					errorMessage: "unable to unmarshal json: invalid character 'n' looking for beginning of object key string",
				},
			},
		},
	}

	for _, test := range cases {
		result, err := getJokes([]byte(test.input))

		// assertions for scenarios that don't produce errors
		if err == nil {
			assert.Len(t, result, len(test.expect.result.slice))
			assert.Equal(t, test.expect.result.slice, result)
		}

		// assertions for scenarios that produce errors
		if err != nil {
			assert.EqualError(t, err, test.expect.result.errorMessage)
			assert.Equal(t, test.expect.result.slice, result)
		}
	}
}

func TestSelectDadJoke(t *testing.T) {
	cases := []struct {
		name  string
		input []string
	}{
		{
			name: "given 3 jokes, selectDadJoke will randomly pick a joke",
			input: []string{
				"option 1",
				"option 2",
				"option 3",
			},
		},
	}

	for _, test := range cases {
		result := selectDadJoke(test.input)

		// testing for probability i.e. does our result show up at least once? This is more of a "transparency" test
		var resultTally []string
		for i := 0; i < 1000; i++ {
			resultTally = append(resultTally, selectDadJoke(test.input))
		}

		assert.Contains(t, resultTally, result)
	}
}
