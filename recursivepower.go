package piscine

func RecursivePower(nb int, power int) int {
	ininb := nb
	if power < 0 {
		return (0)
	}
	if power == 0 {
		return (1)
	}
	for power > 1 {
		nb = nb * ininb
		power--
	}
	return (nb)
}
