package main

import "fmt"

func main() {
	fmt.Print(wordFlip("First second last"))
	fmt.Print(wordFlip(""))
	fmt.Print(wordFlip("     "))
	fmt.Print(wordFlip(" hello  all  of  you! "))
}

func wordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}
	slices := split(str, " ")

	flipSlice := make([]string, 0)
	for _, word := range slices {
		flipSlice = append([]string{word}, flipSlice...)
	}
	joinStr := join(flipSlice)
	return joinStr + "\n"
}

func split(str, sep string) []string {
	var slices []string

	index := 0
	for i := 0; i+len(sep) <= len(str); i++ {
		if str[i:i+len(sep)] == sep {
			if index < i {
				slices = append(slices, str[index:i])
			}
			index = i + len(sep)
		}
	}
	if index < len(str) {
		slices = append(slices, str[index:])
	}
	return slices
}

func join(slices []string) string {
	result := ""
	for i, word := range slices {
		if i != len(slices)-1 {
			result += word + " "
		} else {
			result += word
		}
	}
	return result
}
