/*
	**Instructions**

	Write a function ListClear that deletes all nodes from a linked list l.

    	- Tip: assign the list's pointer to nil.

*/

package linkedlist

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		next := l.Head.Next
		current.Next = nil
		current = next
	}
	l.Head = nil
	l.Tail = nil
}


/*
	It is good practice to also ensure that all nodes are properly dereferenced. In Go, garbage collection handles the actual memory freeing, but it's still a good habit to clean up all references.
*/

/*
	In this implementation, we'll iterate through the list and set each node's Next pointer to nil before moving to the next node. Finally, we will set both Head and Tail to nil to completely clear the list.
*/