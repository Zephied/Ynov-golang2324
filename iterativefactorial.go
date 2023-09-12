package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb < 1 || nb > 12 {
		return (0)
	}
	for nb > 0 {
		n = n * nb
		nb--
	}
	return (n)
}
