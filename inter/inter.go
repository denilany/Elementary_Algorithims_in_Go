package main

import "fmt"

func inter(str1, str2 string) string {
	map1 := make(map[rune]bool)
	result := make([]rune, 0, len(str1))

	for _, ch := range str1 {
		if !map1[ch] && containsRune(str2, ch) {
			map1[ch] = true
			result = append(result, ch)
		}
	}
	return string(result)
}

func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

func main() {
	str1 := "padinton"
	str2 := "paqefwtdjetyiytjneytjoeyjnejeyj"

	result := inter(str1, str2)
	fmt.Println(result)
}
