package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb > 25 || nb < 0 {
		return (0)
	}
	if nb == 0 {
		return (1)
	}
	for nb > 0 {
		n = n * nb
		nb--
	}
	return (n)
}
