package main

import (
	"github.com/01-edu/z01"
)

func PrintRevComb() {
	next := false
	for i := '9'; i >= '2'; i-- {
		for j := i-1; j >= '1'; j-- {
			for k := j-1; k >= '0'; k-- {
				if next {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				next = true

				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintRevComb()
}
