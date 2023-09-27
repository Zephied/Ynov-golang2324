package piscine

func ListSize(l *List) int {
	n := 0
	for l.Head != nil {
		n++
		l.Head = l.Head.Next
	}
	return n
}
