package piscine

func Sqrt(nb int) int {
	n := 1
	for n*n < nb {
		n++
	}
	if n*n == nb {
		return (n)
	} else {
		return (0)
	}
}
