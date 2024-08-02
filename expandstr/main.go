package main

import "fmt"

func main() {
	str1 := "you   see   it's   easy   to   display    the     same  thing"
	str2 := "   only  it's harder   "

	res1 := expandstr(str1)
	res2 := expandstr(str2)

	fmt.Printf("%q\n", res1)
	
	fmt.Println() 

	fmt.Printf("%q\n", res2)
}

func expandstr(str string) string {
	// split the str along white spaces
	slice := split(str, " ")

	// join the split string with three spaces
	joinedStr := join(slice, "   ")
	return joinedStr
}

func split(s, sep string) []string {
	var slice []string

	index := 0
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			if index < i {
				slice = append(slice, s[index:i])
			}
			index = i + len(sep)
		}
	}
	if index < len(s) {
		slice = append(slice, s[index:])
	}
	return slice
}

func join(slice []string, sep string) string {
	var joinedStr string
	for i, word := range slice {
		joinedStr += word
		if i != len(slice)-1 {
			joinedStr += sep
		}
	}
	return joinedStr
}

// func main() {
// 	args := [][]string{
// 		{"hello you"},
// 		{"   only  it's harder   "},
// 		{"you   see   it's   easy   to   display    the     same  thing"},
// 	}

// 	args = append(args, random.StrSlice(chars.Words))

// 	for _, v := range args {
// 		challenge.Program("expandstr", v...)
// 	}
// 	challenge.Program("expandstr")
// }
