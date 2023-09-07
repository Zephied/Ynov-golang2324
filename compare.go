package piscine

func Compare(a, b string) int {
	i := 0
	y := 0
	for range a {
		i++
	}
	for range b {
		y++
	}
	if i != y {
		z := 0
		if a[z] != b[z] {
			return (-1)
		} else {
			return (1)
		}
	}
	return (0)
}
