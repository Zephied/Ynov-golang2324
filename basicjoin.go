package piscine

func BasicJoin(elems []string) string {
	r := ""
	i := 0
	for range elems {
		r += string(elems[i])
		i++
	}
	return (r)
}
