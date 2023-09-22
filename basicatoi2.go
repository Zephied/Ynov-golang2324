package piscine

func BasicAtoi2(s string) int {
	var res int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return res
}
