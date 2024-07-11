package main

import (
	"os"
	"fmt"
)

func binary(num int) string {
	var bin []rune

	for i := 7; i >= 0; i-- {
		bit := (num>>i) & 1
		bin = append(bin, rune(bit+'0'))
	}
	return string(bin)
}

func atoi(str string) int {
	var num int
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			num = num * 10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return num
}

func itoa(num int) string {
	var digits []rune

	for num > 0 {
		digit := num % 10
		digits = append([]rune{rune('0'+digit)}, digits...)
		num /= 10
	}
	return string(digits)
}

func main() {
	args := os.Args[1]
	num := atoi(args) 
	if num == 0 {
		os.Stdout.WriteString("00000000\n")
		return
	}

	str := binary(num)
	
	fmt.Println(str)
}