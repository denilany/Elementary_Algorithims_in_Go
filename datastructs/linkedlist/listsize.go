/*
	**Instructions**

	Write a function ListSize that returns the number of elements in a linked list l.
*/

package linkedlist

func ListSize(l *List) int {
	size := 0
	for current := l.Head; current != nil; current = current.Next {
		size++
	}
	return size
}
