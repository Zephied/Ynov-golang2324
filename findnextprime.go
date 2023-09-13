package piscine

func FindNextPrime(nb int) int {
	i := nb
	if nb%2 == 0 || nb <= 2 {
		return (2)
	}
	for i <= nb/i {
		if nb%i == 0 {
			return (i)
		}
		i++
	}
	return (i)
}
