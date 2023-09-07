package piscine

func AlphaCount(s string) int {
	i := 0
	c := 0
	for range s {
		if s[c] >= 'a' && s[c] <= 'z' || s[c] >= 'A' && s[c] <= 'Z' {
			i++
		}
		c++
	}
	return (i)
}
