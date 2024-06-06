package main

import "fmt"

func activeBits(num int) int {
	count := 0

	for i := 7; i >= 0; i-- {
		bit := (num >> i) & 1
		if num < bit {
			count++
		}
	}
	return count
}

func main() {
	num := 7
	fmt.Println(activeBits(num))
}
