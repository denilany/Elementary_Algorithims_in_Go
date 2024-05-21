package main

import (
	"os"

	"github.com/01-edu/z01"
)

func RotArgs(slice []string) []string {
	length := len(slice)
	if length == 0 {
		return slice
	}

	first := slice[0]
	for i := 0; i < length-1; i++ {
		slice[i] = slice[i+1]
	}
	slice[length-1] = first

	return slice
}

func println(slice []string) {
	length := len(slice)

	for i, str := range slice {
		for _, ch := range str {
			z01.PrintRune(ch)
		}

		if i != length-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) <= 2 {
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1:]
	newSlice := RotArgs(args)

	println(newSlice)
}
