package easy

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	head := linkedList
	current := linkedList
	for current.Next != nil {
		dummy := current.Next
		for dummy != nil && current.Value == dummy.Value {
			dummy = dummy.Next
		}

		current.Next = dummy
		current = current.Next
	}
	return head
}
