package main

import "fmt"

func canJump(uSlice []uint) bool {
	i := 0
	for i < len(uSlice) {
		if i == len(uSlice)-1 {
			return true
		}

		if int(uSlice[i]) == 0 {
			break
		}

		i += int(uSlice[i])
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(canJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(canJump(input2))

	input3 := []uint{0}
	fmt.Println(canJump(input3))

	input4 := []uint{1, 1, 3, 1, 0}
	fmt.Println(canJump(input4))

	input5 := []uint{5, 4, 3, 2, 1, 0}
	fmt.Println(canJump(input5))

	input6 := []uint{1, 2, 3, 0, 2}
	fmt.Println(canJump(input6))
}
