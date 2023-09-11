package piscine

func Join(strs []string, sep string) string {
	r := ""
	i := 0
	for range strs {
		r += string(strs[i])
		if i != len(strs)-1 {
			r += string(sep)
		}
		i++
	}
	return (r)
}
