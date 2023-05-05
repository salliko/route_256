package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertStringToIntArray(value string) []int {
	listString := strings.Split(value, " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}

	return arr
}

func TestTreeRoot(t *testing.T) {
	tests := []struct {
		msg      []int
		expected int
	}{
		{
			msg:      convertStringToIntArray("3 0 1 0 5 2 2 6 4 3 5 1 3 2 0 6 0"),
			expected: 4,
		},
		{
			msg:      convertStringToIntArray("3 1 1 1 2 4 2 4 0 2 0"),
			expected: 3,
		},
		{
			msg:      convertStringToIntArray("1 0 2 1 1"),
			expected: 2,
		},
		{
			msg:      convertStringToIntArray("1 0"),
			expected: 1,
		},
		{
			msg:      convertStringToIntArray("2 1 3 1 0 4 1 2 3 1 1"),
			expected: 4,
		},
		{
			msg:      convertStringToIntArray("4 1 5 5 2 2 1 2 1 3 1 0 3 0"),
			expected: 4,
		},
		{
			msg:      convertStringToIntArray("1 1 2 2 0"),
			expected: 1,
		},
	}

	for _, tt := range tests {
		got := TreeRoot(tt.msg)
		assert.Equal(t, tt.expected, got)
	}
}
