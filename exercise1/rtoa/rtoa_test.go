package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPickLargestGo(t *testing.T) {

	tests := []struct {
		input string
		output int
	}{
		{
			"XLIV",
			44,
		},
		{
			"IX",
			9,
		},
		{
			"C",
			100,
		},
		{
			"I",
			1,
		},
		{
			"",
			0,
		},
		{
			"CLV",
			155,
		},
		{
			"MMMMMMMMMM",
			10000,
		},
		{
			"MMMMMMMMMMI",
			10001,
		},
		{
			"MI",
			1001,
		},

	}

	for _, test := range tests {
		output := pickLargestGo(test.input)
		assert.Equal(t, test.output, output, "invalid value")
	}

}