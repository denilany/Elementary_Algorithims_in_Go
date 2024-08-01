package main

import (
	// student "student"

	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"aaaa",
		"zzzzzZZZZZZ",
		"",
		"ziiiiipMeee",
		"hhellloTthereYouuunggFelllas",
	}

	for _, arg := range args {
		challenge.Function("ZipString", ZipString, solutions.ZipString, arg)
	}
}

func countDupli(s string, i int) int {
	var count int

	for _, ch := range s[i:] {
		if ch == rune(s[i]) {
			count++
		} else {
			break
		}
	}
	return count
}

func ZipString(s string) string {
	result := ""

	i := 0
	for i < len(s) {
		count := countDupli(s, i)
		result += strconv.Itoa(count) + string(s[i])
		i += count
	}
	return result
}

// func main() {
// 	fmt.Println(ZipString("YouuungFellllas"))
// 	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
// 	fmt.Println(ZipString("Helloo Therre!"))
// }
