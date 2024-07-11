package main

func Gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
}

func atoi(str string) int {
	result := 0 

	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		}
	}
	return result
}