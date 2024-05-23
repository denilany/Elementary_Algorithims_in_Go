package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	num := atoi(arg)
	if num < 0 {
		return
	}
	var status string
	if isPowerOf2(num) {
		status = "true"
	} else {
		status = "false"
	}

	printLn(status)
}

func atoi(s string) int {
	var result int
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0
		}
		result = result*10 + int(digit-'0')
	}
	return result
}

// func isPowerOf2(n int) bool {
// 	for n > 1 {
// 		if n%2 != 0 {
// 			return false
// 		}
// 		n /= 2
// 	}
// 	return true
// }

// Using bitwise AND operator (&)
func isPowerOf2(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}

// func itoa(num int) string {
// 	var str string
// 	for num > 0 {
// 		val := num % 10
// 		str += string('0' + val)
// 		num /= 10
// 	}
// 	return str
// }

func printLn(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
