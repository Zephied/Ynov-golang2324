package piscine

func ToLower(s string) string {
	i := 0
	r := ""
	for range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			r += string(s[i] + 32)
		} else {
			r += string(s[i])
		}
		i++
	}
	return (r)
}
