package main

import "fmt"

func atoi(s string) int {
	result := 0
	sign := 1

	if s[0] == '-' {
		sign *= -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	}

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}

	return result * sign
}

func main() {
	fmt.Println(atoi("12345"))
	fmt.Println(atoi("0000000012345"))
	fmt.Println(atoi("012 345"))
	fmt.Println(atoi("Hello World!"))
	fmt.Println(atoi("+1234"))
	fmt.Println(atoi("-1234"))
	fmt.Println(atoi("++1234"))
	fmt.Println(atoi("--1234"))
}
