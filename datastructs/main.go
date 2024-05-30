package main

import (
	"fmt"

	"Algorithims/datastructs/linkedlist"
)

/*
func main() {
	link := &linkedlist.List{}

	linkedlist.ListPushBack(link, "Hello")
	linkedlist.ListPushBack(link, "man")
	linkedlist.ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Print(link.Head.Data, " -> ")
		link.Head = link.Head.Next
	}
	fmt.Println("nil")
}
*/

func main() {
	link := &linkedlist.List{}

	linkedlist.ListPushFront(link, "Hello")
	linkedlist.ListPushFront(link, "man")
	linkedlist.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
