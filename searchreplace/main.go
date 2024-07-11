package main

import(
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	args := os.Args[1:]
	result := searchReplace(args[0], args[1], args[2])
	printLn(result)
}

func searchReplace(str string, letter string, replace string) string {
	result := ""
	for _, ch := range str {
		if string(ch) == letter {
			result += replace
		} else {
			result += string(ch)
		}
	}
	return result
}

func printLn(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
