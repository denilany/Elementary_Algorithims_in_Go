/*
	**Instructions**

	Write a function ListPushBack that inserts a new element NodeL at the end of the list l while using the structure List.

*/

package linkedlist

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}
