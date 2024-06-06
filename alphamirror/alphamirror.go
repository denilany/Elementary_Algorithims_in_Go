package main

import (
	"os"

	"github.com/01-edu/z01"
)

func alphaMirror(str string) string {
	result := ""
	for _, ch := range str {
		if isAlphabet(ch) {
			if ch >= 'a' && ch <= 'z' {
				result += string('z' - (ch - 'a'))
			} else {
				result += string('Z' - (ch - 'A'))
			}
		} else {
			result += string(ch)
		}

	}
	return result
}

func isAlphabet(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func printLn(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	result := alphaMirror(args[0])
	printLn(result)
}
