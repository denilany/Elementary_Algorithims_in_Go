package main

import (
	"fmt"
	"os"
)

func rot13(str string) string {
	result := ""

	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch-'a'+13)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch-'A'+13)%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	res := rot13(args)
	fmt.Println(res)
}
