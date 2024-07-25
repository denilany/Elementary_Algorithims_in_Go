package main

import "fmt"

// func RevConcatAlternate(slice1, slice2 []int) []int {
// 	var result []int

// 	if len(slice1) < len(slice2) {
// 		slice1, slice2 = slice2, slice1
// 	}

// 	for i := len(slice1) - 1; i >= 0; i-- {
// 		result = append(result, slice1[i])
// 		if i < len(slice2) {
// 			result = append(result, slice2[i])
// 		}
// 	}
// 	return result
// }

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)
	maxLen := len1
	if len2 > maxLen {
		maxLen = len2
	}

	result := make([]int, 0, len1+len2)

	for i := maxLen; i >= 0; i-- {
		if i < len1 {
			result = append(result, slice1[i])
		}
		if i < len2 {
			result = append(result, slice2[i])
		}
	}
	return result
}
