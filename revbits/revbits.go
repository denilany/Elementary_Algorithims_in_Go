package main

import "fmt"

func reverseBits(oct byte) byte {
	var reverse byte

	for i := 0; i < 8; i++ {
		bit := (oct >> i) & 1
		reverse = reverse | (bit << (7 - i))
	}
	return reverse
}

func main() {
	var b byte = 0b00100110
	fmt.Printf("%08b", reverseBits(b))
	fmt.Println()
}
