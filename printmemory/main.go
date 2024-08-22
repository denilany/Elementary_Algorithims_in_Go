package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func printstr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func printHex(b byte) string {
	hex := "0123456789abcdef"
	return string(hex[b>>4]) + string(hex[b&0x0f])
}

func PrintMemory(arr [10]byte) {
	// Print hex values
	for i, b := range arr {
		hex := printHex(b)
		printstr(hex)

		if (i+1)%4 == 0 || i == len(arr)-1 {
			printstr("\n")
		} else {
			printstr(" ")
		}
	}

	// Print character representation
	for _, b := range arr {
		if b > 32 && b <= 126 {
			printstr(string(b))
		} else {
			printstr(".")
		}
	}
	printstr("\n")
}

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)
}

// func main() {
// 	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }
