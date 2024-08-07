package main

import (
	"algo/datastructs/linkedlist"
	"fmt"
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

/*
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
*/

/*
func main() {
	link := &linkedlist.List{}

	linkedlist.ListPushFront(link, "Hello")
	linkedlist.ListPushFront(link, "2")
	linkedlist.ListPushFront(link, "you")
	linkedlist.ListPushFront(link, "man")

	fmt.Println(linkedlist.ListSize(link))
}
*/

/*
func main() {
	link := &linkedlist.List{}
	link2 := &linkedlist.List{}

	linkedlist.ListPushBack(link, "three")
	linkedlist.ListPushBack(link, 3)
	linkedlist.ListPushBack(link, "1")

	fmt.Println(linkedlist.ListLast(link))
	fmt.Println(linkedlist.ListLast(link2))
}
*/

/*
type (
	List = linkedlist.List
	Node = linkedlist.NodeL
)

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	linkedlist.ListPushBack(link, "I")
	linkedlist.ListPushBack(link, 1)
	linkedlist.ListPushBack(link, "something")
	linkedlist.ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	linkedlist.ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}
*/

/*
func main() {
	link := &linkedlist.List{}

	linkedlist.ListPushBack(link, "hello")
	linkedlist.ListPushBack(link, "how are")
	linkedlist.ListPushBack(link, "you")
	linkedlist.ListPushBack(link, 1)

	fmt.Println(linkedlist.ListAt(link.Head, 3).Data)
	fmt.Println(linkedlist.ListAt(link.Head, 1).Data)
	fmt.Println(linkedlist.ListAt(link.Head, 7))
}
*/

func main() {
	link := &linkedlist.List{}

	linkedlist.ListPushBack(link, 1)
	linkedlist.ListPushBack(link, 2)
	linkedlist.ListPushBack(link, 3)
	linkedlist.ListPushBack(link, 4)

	linkedlist.ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
