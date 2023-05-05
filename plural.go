package main

import (
	"fmt"
	"strings"
)

func hasSuffixY(world string) bool {
	consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z"}
	for _, val := range consonants {
		suffix := fmt.Sprintf("%s%s", val, "y")
		if strings.HasSuffix(world, suffix) {
			return true
		}
	}
	return false
}

func Plural(word string) string {
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "sh") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") {
		word += "es"
	} else if hasSuffixY(word) {
		word = strings.TrimRight(word, "y")
		word += "ies"
	} else {
		word += "s"
	}
	return word
}
