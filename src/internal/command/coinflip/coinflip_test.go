package coinflip

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

func TestCoinFlip(t *testing.T) {
	cases := []struct {
		name   string
		expect []string
	}{
		{
			name:   "when the function runs, return either heads or tails",
			expect: sides,
		},
	}

	for _, test := range cases {
		result := coinFlip()
		assert.Contains(t, test.expect, result)
	}
}

func TestSelectPhrase(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "when the function runs, return a phrase",
			input:  []string{"phrase one", "phrase two", "phrase three"},
			expect: []string{"phrase one", "phrase two", "phrase three"},
		},
	}

	for _, test := range cases {
		result := selectPhrase(test.input)
		assert.Contains(t, test.expect, result)
	}
}

func TestMakePhrase(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			side   string
			phrase string
		}
		expect string
	}{
		{
			name: "when a phrase has a placeholder at the end of the template, it is substituted with the chosen coin side",
			input: struct {
				side   string
				phrase string
			}{
				side:   "Heads",
				phrase: "It's **{{.Side}}**",
			},
			expect: "It's **Heads**",
		},
		{
			name: "when a phrase has a placeholder at the start of the template, it is substituted with the chosen coin side",
			input: struct {
				side   string
				phrase string
			}{
				side:   "Tails",
				phrase: "**{{.Side}}**, you hate to see it",
			},
			expect: "**Tails**, you hate to see it",
		},
		{
			name: "when a phrase has a placeholder in the middle of the template, it is substituted with the chosen coin side",
			input: struct {
				side   string
				phrase string
			}{
				side:   "Tails",
				phrase: "What the **{{.Side}}** doing?",
			},
			expect: "What the **Tails** doing?",
		},
	}

	for _, test := range cases {
		result := makePhrase(test.input.side, test.input.phrase)
		assert.Equal(t, test.expect, result)
	}
}
