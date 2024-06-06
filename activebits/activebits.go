package main

import "fmt"

// Function to count the number of set bits (1s) in the binary representation of an integer
func activeBits(num int) int {
	count := 0

	// Iterate through each bit of the integer
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		if bit == 1 {
			count++
		}
	}
	return count
}

func main() {
	num := 7
	fmt.Println(activeBits(num)) // Output: 3, since 7 in binary is 111

	num = 15
	fmt.Println(activeBits(num)) // Output: 4, since 15 in binary is 1111
}
