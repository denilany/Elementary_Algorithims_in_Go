package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)

	for _, ch := range str1 {
		if !m1[ch] {
			m1[ch] = true
		}
	}

	for _, ch := range str2 {
		if !m2[ch] {
			m2[ch] = true
		}
	}

	uniqueChars := make(map[rune]bool)
	for char := range m1 {
		if !m2[char] {
			uniqueChars[char] = true
		}
	}

	for char := range m2 {
		if !m1[char] {
			uniqueChars[char] = true
		}
	}

	if len(uniqueChars) == 0 {
		return -1
	}
	return len(uniqueChars)
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("26235e5", "4478q92"))
	fmt.Println(WeAreUnique("w^p@@j", "w^p@@j"))
}
