package main

import "fmt"

func prevPrime(num int) int {
	for i := num; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(prevPrime(5))
	fmt.Println(prevPrime(4))
}
