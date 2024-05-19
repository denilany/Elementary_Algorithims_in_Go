package main

import (
	"os"

	"github.com/01-edu/z01"
)

func containsWord(first, second string) bool {
	j := 0
	for i := 0; i < len(second) && j < len(first); i++ {
		if first[j] == second[i] {
			j++
		}
	}

	return j == len(first)
}

func wordMatch(first, second string) string {
	if containsWord(first, second) {
		return first
	}
	return ""
}

func print(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:3]

	if len(os.Args) != 3 {
		return
	}
	first := args[0]
	second := args[1]
	word := wordMatch(first, second)
	print(word)
}
