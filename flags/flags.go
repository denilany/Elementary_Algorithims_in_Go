package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Checks if the input string is a valid insert flag
func isIFlag(s string) bool {
	return strings.HasPrefix(s, "-i=") || strings.HasPrefix(s, "--insert=")
}

// Checks if the input string is a valid order flag
func isOFlag(s string) bool {
	return s == "-o" || s == "--order"
}

// Handles the sorting of the input string
func printInOrder(s string) {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	fmt.Println(string(runes))
}

// Main function to handle command-line arguments and process flags
func main() {
	if len(os.Args) <= 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	var str, insertStr string
	isOrder := false

	for _, arg := range os.Args[1:] {
		if isIFlag(arg) {
			insertStr += getInsertString(arg)
		} else if isOFlag(arg) {
			isOrder = true
		} else {
			str += arg
		}
	}

	str += insertStr
	if isOrder {
		printInOrder(str)
	} else {
		fmt.Println(str)
	}
}

// Prints the help message
func printHelp() {
	fmt.Println("--insert, -i=<string>")
	fmt.Println("\tInsert the specified string into the argument.")
	fmt.Println("--order, -o")
	fmt.Println("\tSort the argument alphabetically.")
}

// Extracts the string to be inserted from the insert flag
func getInsertString(arg string) string {
	parts := strings.SplitN(arg, "=", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
