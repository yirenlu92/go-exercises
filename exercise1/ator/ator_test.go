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
			"XXXXIIII",
		},
		{
			9,
			"VIIII",
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

	}

	for _, test := range tests {
		output := pickLargest(test.input)
		assert.Equal(t, test.output, output, "invalid value")
	}

}