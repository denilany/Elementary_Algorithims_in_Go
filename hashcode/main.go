package main

import "fmt"

func HashCode(dec string) string {
	size := len(dec)
	hashed := make([]rune, size)

	for i, ch := range dec {
		hashedChar := (int(ch) + size) % 127

		if hashedChar < 32 || hashedChar > 126 {
			hashedChar += 33
		}

		hashed[i] = rune(hashedChar)
	}
	return string(hashed)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
