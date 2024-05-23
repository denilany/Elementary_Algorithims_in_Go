package main

import "fmt"

func reverseBits(oct byte) byte {
	var result byte

	for i := 0; i < 8; i++ {
		// Shift the result to the left to make space for the next bit
		result <<= 1

		// Add the rightmost bit of the oct to result
		result |= oct & 1

		// Shift oct to the right to process the next bit
		oct >>= 1
	}
	return result
}

func main() {
	oct := byte(0b00100110)
	reversed := reverseBits(oct)
	fmt.Printf("%08b\n", oct)      // Output: 00100110
	fmt.Printf("%08b\n", reversed) // Output: 01100100
}
