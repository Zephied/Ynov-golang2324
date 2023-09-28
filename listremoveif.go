package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
		return
	}
	previous := l.Head
	current := l.Head.Next
	for current != nil {
		if current.Data == data_ref {
			previous.Next = current.Next
			if current.Next == nil {
				l.Tail = previous
			}
		}
		previous = current
		current = current.Next
	}

}
