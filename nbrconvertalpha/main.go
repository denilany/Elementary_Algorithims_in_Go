package main

import (
	"os"
	"unicode"

	"github.com/01-edu/z01"
)

func atoi(str string) int {
	isNeg := false
	result := 0
	index := 0
	length := len(str)

	// Check for optional leading whitespace and skip it
	for index < length && unicode.IsSpace(rune(str[index])) {
		index++
	}

	// Check for optional sign
	if index < length && (str[index] == '+' || str[index] == '-') {
		if str[index] == '-' {
			isNeg = true
		}
		index++
	}

	// Skip leading zeros
	for index < length && str[index] == '0' {
		index++
	}

	// Convert characters to integer
	for index < length && str[index] >= '0' && str[index] <= '9' {
		result *= 10
		result += int(str[index] - '0')
		index++
	}

	// Apply negative sign if needed
	if isNeg {
		result *= -1
	}

	// Check if there are non-numeric characters left in the string
	if index < length && !unicode.IsSpace(rune(str[index])) {
		return 0
	}

	return result
}

func main() {
	args := os.Args[1:]

	if len(args) > 1 && args[0] == "--upper" {
		for i := 1; i < len(args); i++ {
			num := atoi(args[i])
			if num >= 1 && num <= 26 {
				z01.PrintRune(rune(num - 1 + 'A'))
			} else {
				z01.PrintRune(' ')
			}
		}
	} else {
		for i := 0; i < len(args); i++ {
			num := atoi(args[i])
			if num >= 1 && num <= 26 {
				z01.PrintRune(rune(num - 1 + 'a'))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	if len(os.Args) > 1 {
		z01.PrintRune('\n')
	}
}
