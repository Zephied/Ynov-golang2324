package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	p := l.Head
	for p != nil {
		if comp(p.Data, ref) {
			return &p.Data
		}
		p = p.Next
	}
	return nil
}
