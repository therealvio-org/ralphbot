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
