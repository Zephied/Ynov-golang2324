package piscine

func IsPrime(nb int) bool {
	i := 3
	n := nb
	if n == 2 {
		return (true)
	}
	if n%2 == 0 || n < 2 {
		return (false)
	}
	for i <= n/i {
		if n%i == 0 {
			return (false)
		}
		i++
	}
	return (true)
}
