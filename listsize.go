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
	n := 1
	for l.Head.Next != nil {
		n++
		l.Head = l.Head.Next
	}
	return n
}
