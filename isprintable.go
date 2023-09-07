package piscine

func IsPrintable(s string) bool {
	c := 0
	for range s {
		if s[c] >= 32 && s[c] <= 126 {
			c++
		} else {
			return (false)
		}

	}
	return (true)
}
