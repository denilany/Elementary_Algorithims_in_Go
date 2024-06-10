package main

import (
	"fmt"
)

func rot13(str string) string {
	rotStr := ""

	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			rotStr += string((ch-'A'+13)%26 + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			rotStr += string((ch-'a'+13)%26 + 'a')
		} else {
			rotStr += string(ch)
		}
	}
	return rotStr
}

func main() {
	str := "Hello@ 123 World"
	rot := rot13(str)
	fmt.Println(rot)
}
