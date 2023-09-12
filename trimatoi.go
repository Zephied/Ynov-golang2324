package piscine

func TrimAtoi(s string) int {
	signe := 1
	val := 0
	for _, va := range s {
		if va >= '0' && va <= '9' {
			val = val*10 + int(va-'0')
		} else if va == '-' && val == 0 {
			signe = -1
		}
	}
	return val * signe
}
