package main

import (
	"github.com/01-edu/z01"
)

func lastRune(s string) rune {
	runes := []rune(s)
	last := len(runes) - 1
	return runes[last]
}

func main() {
	z01.PrintRune(lastRune("Hello!"))
	z01.PrintRune(lastRune("Salut!"))
	z01.PrintRune(lastRune("Ola!"))
	z01.PrintRune('\n')
}
