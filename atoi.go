package piscine

func Atoi(s string) int {
	val := 0
	var signe int
	for _, va := range s {
		if va == '-' || va == '+' {
			signe++
		}
		if signe > 1 {
			return 0
		}
	}
	signe = 1
	for _, va := range s {
		if va >= '0' && va <= '9' {
			val = val*10 + int(va-'0')
		} else if va == '-' {
			signe = -1
		} else if va == '+' {
			signe = 1
		} else {
			return 0
		}
	}
	return val * signe
}
