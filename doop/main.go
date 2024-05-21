package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "/" || s == "*" || s == "%"
}

func isVal(s string) bool {
	for _, val := range s {
		if val < '0' || val > '9' {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	result := 0
	sign := 1

	for i, ch := range s {
		if i == 0 && ch == '-' {
			sign = -1
			continue
		} else if i == 0 && ch == '+' {
			continue
		}
		if ch < '0' && ch > '9' {
			return 0
		}
		result *= 10
		result += int(ch - '0')
	}
	return result * sign
}

func itoa(n int) string {
	// If n is zero
	if n == 0 {
		return "0"
	}

	// Negative number
	isNeg := false
	if n < 0 {
		isNeg = true
		n = -n
	}

	// Convert to string
	var result []byte
	for n > 0 {
		digit := n % 10
		result = append([]byte{byte('0' + digit)}, result...)
		n /= 10
	}
	if isNeg {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func println(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	// z01.PrintRune('\n')
}

func doop(val1, val2, sign string) int {
	num1 := atoi(val1)
	num2 := atoi(val2)

	result := 0

	switch sign {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		output := "No division by "
		if num2 == 0 {
			println(output)
			break
		}
		result = num1 / num2

	case "%":
		str := "No modulo by "
		if num2 == 0 {
			println(str)
			break
		}
		result = num1 % num2
	}
	return result
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	args := os.Args[1:]
	// fmt.Println(len(args))

	val1 := args[0]
	val2 := args[2]
	sign := args[1]

	if !isVal(val1) || !isVal(val2) {
		return
	}

	if !isOperator(sign) {
		return
	}

	result := doop(val1, val2, sign)

	str := itoa(result)

	println(str+"\n")
	z01.PrintRune('\n')
}
