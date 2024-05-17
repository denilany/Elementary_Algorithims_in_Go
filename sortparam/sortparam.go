package main

import "fmt"

func Sorted(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ { // use this: for j := 0; j < n-1; j++ if it is in descening order
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{2, 4, 1, 7, 5, 3, 6}
	Sorted(arr)
}
