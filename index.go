package piscine

func Index(s string, toFind string) int {
	i := 0
	j := 0
	y := len([]rune(toFind))
	for range s {
		if s[i+j] == toFind[j] && j <= y {
			j++
		} else {
			i++
			j = 0
		}
		if j == y {
			return (i)
		}
	}
	return (-1)
}
