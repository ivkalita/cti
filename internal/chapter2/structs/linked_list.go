package structs

// LinkedList – single linked list
type LinkedList struct {
	Head *Node
	len  int
}

// Node – node of a single linked list
type Node struct {
	Data int
	Next *Node
}

// NewLinkedList – makes a new list
func NewLinkedList(in []int) LinkedList {
	l := LinkedList{}
	for _, v := range in {
		l.Append(v)
	}

	return l
}

// Append – adds a new value to the end of the list
func (l *LinkedList) Append(value int) {
	l.len++

	n := Node{Data: value}

	if l.Head == nil {
		l.Head = &n
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &n
}

// ToSlice – returns an int slice
func (l LinkedList) ToSlice() []int {
	result := make([]int, 0, l.len)
	cur := l.Head
	for cur != nil {
		result = append(result, cur.Data)
		cur = cur.Next
	}

	return result
}
