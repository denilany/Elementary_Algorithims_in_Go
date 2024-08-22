package main

import (
	"fmt"
)

func printhex(d byte) string {
	hexaDec := "0123456789abcdef"
	return string(hexaDec[d>>4]) + string(hexaDec[d&0x0f])
}

func main() {
	table := []int{16, 42, 56, 75}

	for _, val := range table {
		hex := printhex(byte(val))
		fmt.Println(hex)
	}
}
