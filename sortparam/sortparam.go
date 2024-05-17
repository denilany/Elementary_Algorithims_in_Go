package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Sorted(arr []string) []string {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ { // use this: (for j := 0; j < n-1; j++), if it is in descending order
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func print(arr []string) {
	for _, str := range arr {
		z01.PrintRune(rune(str[0]))
		z01.PrintRune('\n')
	}
}

func main() {
	args := os.Args[1:]
	print(Sorted(args))
}
