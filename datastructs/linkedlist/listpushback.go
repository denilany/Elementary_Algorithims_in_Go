package linkedlist

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func LinkedList(l *List, data interface{}) {
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
