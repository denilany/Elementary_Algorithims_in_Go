package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	// splitMath := strings.Split(romanMath, "+")
	// for _, ch := range splitMath {
	// 	if len(ch) > 1 {
	// 		for _, rn := range ch {

	// 		}
	// 	}
	// }

	return romanStr, romanMath
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, formula := rn(num)

	fmt.Println(formula)
	fmt.Println(roman)
}
