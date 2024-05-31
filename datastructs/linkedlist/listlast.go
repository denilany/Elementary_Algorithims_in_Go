/*
	**Instructions**

	Write a function ListLast that returns the Data interface of the last element of a linked list l.
*/

package linkedlist

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
