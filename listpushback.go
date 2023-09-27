package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = n
	} else {
		l.Head.Next = n
	}
	l.Head = n
}
