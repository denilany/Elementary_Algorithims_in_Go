package main

import (
	"os"

	"github.com/01-edu/z01"
)

func lastWord(str string) string {
	slice := split(str, " ")

	for i := len(slice) - 1; i >= 0; i++ {
		if slice[i] != "" {
			return slice[i]
		}
	}
	return ""
}

func split(s, sep string) []string {
	result := make([]string, 0)

	startIndex := 0
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			if startIndex < i { // Ensure we don't add empty strings
				result = append(result, s[startIndex:i])
			}
			startIndex = i + len(sep)
		}
	}
	if startIndex < len(s) {
		result = append(result, s[startIndex:])
	}
	return result
}

func printLn(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args

	if len(args) != 2 {
		return
	}
	result := lastWord(args[1])

	if result != "" {
		printLn(result)
	}
}
