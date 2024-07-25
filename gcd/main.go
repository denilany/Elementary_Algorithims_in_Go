package main

import "fmt"

func Gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// func atoi(str string) int {
// 	result := 0

// 	for _, ch := range str {
// 		if ch >= '0' && ch <= '9' {
// 			result = result*10 + int(ch-'0')
// 		}
// 	}
// 	return result
// }

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
