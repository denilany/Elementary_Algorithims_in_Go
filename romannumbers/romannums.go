package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	arg := os.Args[1]
	num := atoi(arg)
	if num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, formula := rn(num)

	printLn(formula)
	printLn(roman)
}

func romanMap() map[int]string {
	return map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
}

func rn(num int) (string, string) {
	romanStr := ""
	romanMath := ""
	romanNumerals := romanMap()

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range values {
		for num >= value {
			num -= value
			romanStr += romanNumerals[value]
			if romanMath != "" {
				romanMath += "+"
			}
			romanMath += romanNumerals[value]
		}
	}
	splitMath := split(romanMath, "+")
	for i, ch := range splitMath {
		if len(ch) > 1 {
			ch = "(" + string(ch[1]) + "-" + string(ch[0]) + ")"
			splitMath = append(splitMath[:i], ch)
		}
	}
	joined := join(splitMath, "+")

	return romanStr, joined
}

func printLn(val string) {
	for _, ch := range val {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func split(s, sep string) []string {
	var result []string

	index := 0
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			if index < i {
				result = append(result, s[index:i])
			}
			index = i + len(sep)
		}
	}
	if index < len(s) {
		result = append(result, s[index:])
	}
	return result
}

func join(arr []string, sep string) string {
	result := ""
	for i, s := range arr {
		if i != len(arr)-1 {
			result += s + sep
		} else {
			result += s
		}
	}
	return result
}

func atoi(str string) int {

	result := 0
	sign := 1

	for i, ch := range str {
		if ch < '0' || ch > '9' {
			return 0
		}
		if i == 0 && ch == '-' {
			sign = -1
			continue
		} else if i == 0 && ch == '+' {
			continue
		}
		result *= 10
		result += int(ch - '0')
	}
	return result * sign
}
