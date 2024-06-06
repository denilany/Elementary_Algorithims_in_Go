package main

import "fmt"

func reduceInt(a []int, f func(int, int) int) int {
	if len(a) == 0 {
		return 0
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	return result
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 50, 40, 20}
	fmt.Println("mult: ", reduceInt(as, mul))
	fmt.Println("sum: ", reduceInt(as, sum))
	fmt.Println("div", reduceInt(as, div))
}
