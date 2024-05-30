package main

import (
	"fmt"

	"Algorithims/datastructs/linkedlist"
)

func main() {
	link := &linkedlist.List{}

	linkedlist.LinkedList(link, "Hello")
	linkedlist.LinkedList(link, "man")
	linkedlist.LinkedList(link, "how are you")

	for link.Head != nil {
		fmt.Print(link.Head.Data, " -> ")
		link.Head = link.Head.Next
	}
	fmt.Println("nil")
}
