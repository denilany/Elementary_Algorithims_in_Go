package main

import (
	"fmt"
)

func split(s, sep string) []string {
	result := make([]string, 0)

	startIndex := 0
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			if startIndex < i {
				result = append(result, s[startIndex:i])
			}
			startIndex = i + len(sep)
		}
	}
	// Append the last part of the string if it is present
	if startIndex < len(s) {
		result = append(result, s[startIndex:])
	}
	return result
}

func main() {
	s := "HelloHAPhowHAPareHAPyou?"
	fmt.Printf("%#v\n", split(s, "HAP"))
}
