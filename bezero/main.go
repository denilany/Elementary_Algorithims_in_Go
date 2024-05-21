package main

import "fmt"

func BeZero(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}

	newSlice := make([]int, len(slice))
	for i := range slice {
		newSlice[i] = 0
	}
	return newSlice
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	newArr := BeZero(arr)
	fmt.Println(newArr)
}
