package piscine

func IterativeFactorial(nb int) int {
	n := 1
	for nb > 0 {
		n = n * nb
		nb--
	}
	return (n)
}
