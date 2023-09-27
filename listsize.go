package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	n := 0
	for l.Head != nil {
		n++
		l.Head = l.Head.Next
	}
	return n
}
