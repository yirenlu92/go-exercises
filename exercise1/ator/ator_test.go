package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickLargest(t *testing.T) {

	tests := []struct {
		input int
		output string
	}{
		{
			44,
			"XLIV",
		},
		{
			9,
			"IX",
		},
		{
			100,
			"C",
		},
		{
			1,
			"I",
		},
		{
			0,
			"",
		},
		{
			155,
			"CLV",
		},
		{
			10000,
			"MMMMMMMMMM",
		},
		{
			10001,
			"MMMMMMMMMMI",
		},
		{
			1001,
			"MI",
		},

	}

	for _, test := range tests {
		output := pickLargest(test.input)
		assert.Equal(t, test.output, output, "invalid value")
	}

}