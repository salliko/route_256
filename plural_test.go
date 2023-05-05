package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlural(t *testing.T) {
	tests := []struct {
		msg      string
		expected string
	}{
		{
			msg:      "cat",
			expected: "cats",
		},
		{
			msg:      "tax",
			expected: "taxes",
		},
		{
			msg:      "city",
			expected: "cities",
		},
		{
			msg:      "roof",
			expected: "roofs",
		},
		{
			msg:      "chief",
			expected: "chiefs",
		},
		{
			msg:      "bus",
			expected: "buses",
		},
		{
			msg:      "blitz",
			expected: "blitzes",
		},
		{
			msg:      "marsh",
			expected: "marshes",
		},
		{
			msg:      "house",
			expected: "houses",
		},
		{
			msg:      "puppy",
			expected: "puppies",
		},
		{
			msg:      "counterpunch",
			expected: "counterpunches",
		},
		{
			msg:      "array",
			expected: "arrays",
		},
	}

	for _, tt := range tests {
		got := Plural(tt.msg)
		assert.Equal(t, tt.expected, got)
	}
}
