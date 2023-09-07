package piscine

func AlphaCount(s string) int {
	i := 0
	c := 0
	for range s {
		if s[c] >= '0' && s[c] <= '9' {
			i++
		}
		c++
	}
	return (i)
}
