package piscine

func Capitalize(s string) string {
	i := 0
	r := ""
	for range s {
		if i == 0 && s[i] >= 'a' && s[i] <= 'z' {
			r += string(s[i] - 32)
			i++
		} else {
			if i == 0 {
				r += string(s[0])
				i++
			} else {
				if s[i-1] < '0' || s[i-1] > '9' && s[i-1] < 'A' || s[i-1] > 'Z' && s[i-1] < 'a' || s[i-1] > 'z' {
					if s[i] >= 'a' && s[i] <= 'z' {
						r += string(s[i] - 32)
					} else {
						r += string(s[i])
					}
					i++
				} else {
					if s[i] >= 'A' && s[i] <= 'Z' {
						r += string(s[i] + 32)
					} else {
						r += string(s[i])
					}
					i++
				}
			}
		}
	}
	return (r)
}
