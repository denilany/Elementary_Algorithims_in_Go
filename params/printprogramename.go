package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	runes := []rune(os.Args[0])
	index := findIndex(runes, '/')
	printStr(args[index:])
	z01.PrintRune('\n')
}

func findIndex(runes []rune, toFind rune) int {
	var index int

	for i, char := range runes {
		if char == toFind {
			index = i + 1
		}
	}
	return index
}

func printStr(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
}
