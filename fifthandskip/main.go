package main

import "fmt"

func FifthAndSkip(s string) string {
	if len([]rune(s)) < 5 {
		return "Invalid Input\n"
	}

	str := removeSpaces(s)

	result := ""

	count := 0
	for _, ch := range str {
		if count == 5 {
			result += " "
			count = 0
			continue
		}
		result += string(ch)
		count++
	}
	return result + "\n"
}

func removeSpaces(str string) string {
	result := ""
	for _, ch := range str {
		if ch != ' ' {
			result += string(ch)
		}
	}
	return result
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	// fmt.Print(FifthAndSkip("1234"))
	fmt.Print(FifthAndSkip("e 5£ @ 8* 7 =56 ;"))
}
