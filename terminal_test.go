package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotepad(t *testing.T) {
	tests := []struct {
		msg      string
		expected string
	}{
		{
			msg:      "otLLLrRuEe256LLLN",
			expected: "route\n256",
		},
		// FAIL
		// {
		// 	msg:      "LRLUUDE",
		// 	expected: "",
		// },
		{
			msg:      "itisatest",
			expected: "itisatest",
		},
		// FAIL
		// {
		// 	msg:      "abNcdLLLeUfNxNx",
		// 	expected: "af\nx\nxb\necd",
		// },
	}

	for _, tt := range tests {
		got := Notepad(tt.msg)
		assert.Equal(t, tt.expected, got)
	}
}
