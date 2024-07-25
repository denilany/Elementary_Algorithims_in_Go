package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	if len(nbrs) == 0 {
		return a
	}

	start := nbrs[0]
	end := length

	if len(nbrs) > 1 {
		end = nbrs[1]
	}
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}
	if end < start {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
