package main

import (
	"fmt"
	"os"
	"strconv"
)

func romanMap(val int) string {
	roman := map[int]string{
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
	return roman[val]
}

func rn(num int) string {
	romanStr := ""

	// Thousands
	thousands := num / 1000
	if thousands < 4000 {
		for i := 0; i < thousands; i++ {
			romanStr += romanMap(1000)
		}
	} else {
		fmt.Println("ERROR: cannot convert to roman digit")
		return ""
	}
	num %= 1000

	//Hundreds
	hundreds := num / 100
	if hundreds < 4 {
		for i := 0; i < hundreds; i++ {
			romanStr += romanMap(100)
		}
	} else if hundreds > 5 && hundreds < 9 {
		romanStr += romanMap(500)
		for i := 0; i < hundreds-5; i++ {
			romanStr += romanMap(100)
		}
	} else {
		romanStr += romanMap(hundreds * 100)
	}
	num %= 100

	//Tens
	tens := num / 10
	// fmt.Println(tens)
	if tens < 4 {
		for i := 0; i < tens; i++ {
			romanStr += romanMap(10)
		}
	} else if tens > 50 && tens < 90 {
		romanStr += romanMap(50)
		for i := 0; i < tens-5; i++ {
			romanStr += romanMap(10)
		}
	} else {
		romanStr += romanMap(tens * 10)
	}
	num %= 10

	//Ones
	ones := num

	if ones < 4 {
		for i := 0; i < ones; i++ {
			romanStr += romanMap(1)
		}
	} else if ones > 5 && ones < 9 {
		romanStr += romanMap(5)
		for i := 0; i < ones-5; i++ {
			romanStr += romanMap(1)
		}
	} else {
		romanStr += romanMap(ones)
	}
	return romanStr
}

func main() {
	args := os.Args[1:]
	val, _ := strconv.Atoi(args[0])

	roman := rn(val)
	fmt.Println(roman)
}
