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
		ListRemoveIf(l, data_ref)
		return
	}
	iterator := l.Head
	for iterator.Next != nil {
		if iterator.Next.Data == data_ref {
			iterator.Next = iterator.Next.Next
		} else {
			iterator = iterator.Next
		}
	}
}
