package piscine

func IterativePower(nb int, power int) int {
	ininb := nb
	if power < 0 {
		return (0)
	}
	for power > 1 {
		nb = nb * ininb
		power--
	}
	return (nb)
}
