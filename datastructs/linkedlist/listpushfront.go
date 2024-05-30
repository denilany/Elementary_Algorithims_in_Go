/*
	**Instructions**

	Write a function ListPushFront that inserts a new element NodeL at the beginning of the list l while using the structure List
*/

package linkedlist

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{
		Data: data,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = node
		return
	} else {
		node.Next = l.Head
		l.Head = node
	}
}
