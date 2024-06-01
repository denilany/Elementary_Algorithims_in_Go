/*
	**Instructions**

	Write a function ListReverse that reverses the order of the elements of a given linked list l.

*/

package linkedlist

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	var next *NodeL

	// Traverse the list and reverse the links
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// After the loop, prev is the new head of the reversed list
    // Set the tail to the original head
	l.Tail = l.Head

	l.Head = prev
}
