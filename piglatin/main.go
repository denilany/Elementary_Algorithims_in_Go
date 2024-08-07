package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	result := pigLatin(os.Args[1])

	if result == os.Args[1] {
		printLn("No vowels")
		return
	}
	printLn(result)
}

func pigLatin(str string) string {
	// nonVowel := ""

	for i, ch := range str {
		if isVowel(ch) {
			if i == 0 {
				return str + "ay"
			}
			return str[i:] + str[:i] + "ay"
		}
	}
	return str
}

func isVowel(vow rune) bool {
	return vow == 'a' || vow == 'e' || vow == 'i' || vow == 'o' || vow == 'u' || vow == 'A' || vow == 'E' || vow == 'I' || vow == 'O' || vow == 'U'
}

func printLn(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
