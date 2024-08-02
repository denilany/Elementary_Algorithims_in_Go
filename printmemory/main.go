package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

// func printstr(str string) {
// 	for _, char := range str {
// 		z01.PrintRune(char)
// 	}
// }

// func hex(b byte) string {
// 	hex := "0123456789abcdef"
// 	return string(hex[b>>4]) + string(hex[b&0x0f])
// }

func PrintMemory(arr [10]byte) {
	// Print hex values
	for i, b := range arr {
		fmt.Printf("%02x", b)
		if (i+1)%4 == 0 && i != len(arr)-1 {
			fmt.Print("\n")
		} else if i != len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")

	// Print character representation
	for _, b := range arr {
		if b > 32 && b <= 126 {
			fmt.Print(string(b))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")
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
