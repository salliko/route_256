package main

import (
	"strings"
	"unicode"
)

func countChars(input []string) int {
	chars := 0
	for _, char := range input {
		if strings.ToUpper(char) != char || unicode.IsNumber([]rune(char)[0]) {
			chars++
		}
	}
	return chars
}

func countRows(input []string) int {
	countRows := 1
	for _, char := range input {
		if char == "N" {
			countRows++
		}
	}
	return countRows
}

func getEndIndex(input []string) int {
	for i, val := range input {
		if val == "" {
			return i
		}
	}
	return len(input) - 1
}

func splitResult(arr [][]string) string {
	var result []string
	for _, raw := range arr {
		result = append(result, strings.Join(raw, ""))
	}

	return strings.Join(result, "\n")
}

func easyTerminal(input []string) [][]string {
	charsLen := countChars(input)
	rowsLen := countRows(input)

	result := make([][]string, rowsLen)
	for i := range result {
		result[i] = make([]string, charsLen)
	}

	cur := []int{0, 0}

	for _, char := range input {
		if strings.ToUpper(char) == char && !unicode.IsNumber([]rune(char)[0]) {
			if char == "L" {
				if cur[1] != 0 {
					cur[1]--
				}
			} else if char == "R" {
				if cur[1] <= len(result[cur[0]]) {
					cur[1]++
				}
			} else if char == "H" {
				cur[1] = 0
			} else if char == "E" {
				cur[1] = getEndIndex(result[cur[0]])
			} else if char == "N" {
				// Подумать над сдвигом строк.
				newRow := cur[0] + 1
				currentIndex := 0
				for i := 0; i < len(result[0]); i++ {
					if i >= cur[1] {
						result[newRow][currentIndex] = result[0][i]
						result[0][i] = ""
						currentIndex++
					}
				}

				cur[0] = newRow
				cur[1] = currentIndex
			} else if char == "U" {
				if cur[0]-1 > 0 {
					cur[0]--
				}
				if result[cur[0]][cur[1]] == "" {
					cur[1] = getEndIndex(result[cur[0]])
				}
			} else if char == "D" {
				if cur[0] < len(result) {
					cur[0]++
				}
				if result[cur[0]][cur[1]] == "" {
					cur[1] = getEndIndex(result[cur[0]])
				}
			}
		} else {
			if result[cur[0]][cur[1]] != "" {
				tmpArr := make([]string, charsLen)
				for idx := 0; idx < len(tmpArr); idx++ {
					if idx < cur[1] {
						tmpArr[idx] = result[cur[0]][idx]
					} else if idx == cur[1] {
						tmpArr[idx] = char
					} else {
						tmpArr[idx] = result[cur[0]][idx-1]
					}
				}
				result[cur[0]] = tmpArr
			} else {
				result[cur[0]][cur[1]] = char
			}
			cur[1]++
		}
	}

	return result
}

func Notepad(input string) string {
	unput_data := strings.Split(input, "")
	result := easyTerminal(unput_data)
	return splitResult(result)
}
