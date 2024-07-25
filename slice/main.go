package main

import "fmt"

// func Slice(a []string, nbrs ...int) []string {
// 	length := len(a)

// 	if len(nbrs) == 0 {
// 		return a
// 	}

// 	start := nbrs[0]
// 	end := length

// 	if len(nbrs) > 1 {
// 		end = nbrs[1]
// 	}
// 	if start < 0 {
// 		start = length + start
// 	}
// 	if end < 0 {
// 		end = length + end
// 	}
// 	if end < start {
// 		return nil
// 	}

// 	return a[start:end]
// }

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

	// Handle negative indices
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// Ensure start and end are within bounds
	if start < 0 {
		// start = 0
		return []string{}
	}
	if end > length {
		// end = length
		return []string{}
	}

	// Return nil if start is greater than end or length
	if start >= length || start >= end {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"UgwqV j kfjyv", "H qFXtrF VQ S", "3 VbCFcghtTnh", "vxpXjd531Ns6H", "E4 JUqjNHn8aS", "4 3ZtL0UG 0dm", "la lCjvhA z S", "onCYQTCoqfadc"}

	fmt.Printf("%#v\n", Slice(a, -18, -13))
}

// func main() {
// 	elems := [][]interface{}{
// 		{
// 			[]string{"coding", "algorithm", "ascii", "package", "golang"},
// 			1,
// 		},
// 		{
// 			[]string{"coding", "algorithm", "ascii", "package", "golang"},
// 			-3,
// 		},
// 		{
// 			[]string{"coding", "algorithm", "ascii", "package", "golang"},
// 			2, 4,
// 		},
// 		{
// 			[]string{"coding", "algorithm", "ascii", "package", "golang"},
// 			-2, -1,
// 		},
// 		{
// 			[]string{"coding", "algorithm", "ascii", "package", "golang"},
// 			2, 0,
// 		},
// 	}

// 	s := random.StrSlice(chars.Words)

// 	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

// 	for i := 0; i < 3; i++ {
// 		s = random.StrSlice(chars.Words)
// 		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
// 	}

// 	for _, a := range elems {
// 		challenge.Function("Slice", Slice, solutions.Slice, a...)
// 	}
// }
