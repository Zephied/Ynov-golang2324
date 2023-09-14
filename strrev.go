package piscine

func StrRev(s string) string {
	r := ""
	i := 0
	for range s {
		i++
	}
	i--
	for i > 0 {
		r += string(s[i])
		i--
	}
	return r
}
