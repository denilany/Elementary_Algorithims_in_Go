package main

import (
	"fmt"
)

func trimSpace(s string) string {
	start := 0
	end := len(s) - 1

	for start <= end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}
	for end >= start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n') {
		end--
	}

	return s[start:end]
}

func main() {
	input := "    word is this   "

	trimmedStr := trimSpace(input)
	fmt.Println(trimmedStr)
}
