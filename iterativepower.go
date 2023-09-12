package piscine

func IterativePower(nb int, power int) int {
	ininb := nb
	for power > 1 {
		nb = nb * ininb
		power--
	}
	return (nb)
}
