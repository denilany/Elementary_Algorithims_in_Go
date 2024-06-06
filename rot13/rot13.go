package main

import "fmt"

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
	str := "abc"

	res := rot13(str)
	fmt.Println(res)
}
