package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	current := l.Head
	var prev *NodeL
	l.Tail = l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
