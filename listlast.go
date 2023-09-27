package piscine

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
