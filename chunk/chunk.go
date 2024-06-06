package main

import "fmt"

func chunk(slice []int, size int) {
	result := make([][]int, 0)
	if size == 0 {
		fmt.Println()
		return
	}

	for start := 0; start < len(slice); start += size {
		end := start + size
		if end > len(slice) {
			end = len(slice)
		}
		
		result = append(result, slice[start:end])
	}
	fmt.Println( result)
}

func main() {
	chunk([]int{}, 10)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
