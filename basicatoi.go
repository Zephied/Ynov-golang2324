package piscine

func BasicAtoi(s string) int {
	val := 0
	for _, va := range s {
		if va >= '0' && va <= '9' {
			val = val*10 + int(va-'0')
		}
	}
	return val
}
