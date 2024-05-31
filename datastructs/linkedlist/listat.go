/*
	Instructions

	Write a function ListAt that takes a pointer to the head of the list l and an int pos as parameters. This function should return the pointer to the NodeL in the position pos of the linked list l.

    	- In case of error the function should return nil.
*/

package linkedlist

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	current := l
	currentIndex := 0

	for current != nil {

		if currentIndex == pos {
			return current
		}
		current = current.Next
		currentIndex++
	}

	return nil
}
