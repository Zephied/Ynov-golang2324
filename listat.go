package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	for i := 0; i < pos; i++ {
		if l == nil {
			return nil
		} else {
			l = l.Next
		}
	}
}
