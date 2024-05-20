package main

import "fmt"

func isSep(str byte) bool {
	return str == ' ' || str == '\n' || str == '\t' || str == '\v'
}

func splitWhiteSpaces(s string) []string {
	result := make([]string, 0)

	length := len(s)

	index := 0
	for i := 0; i+1 < length; i++ {
		if isSep(s[i]) {
			result = append(result, s[index:i])
			index = i + 1
		}
	}
	if index < length {
		result = append(result, s[index:])
	}
	return result
}

func main() {
	s := "Hello  how  are  you?"
	slice := splitWhiteSpaces(s)
	fmt.Printf("%#v", slice)
}
