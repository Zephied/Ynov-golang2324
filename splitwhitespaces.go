package piscine

func SplitWhiteSpaces(s string) []string {
	start := 0
	slr := []string{}
	for i, v := range s {
		if v == ' ' {
			if !(s[i] == ' ' && s[start] == ' ') {
				slr = append(slr, s[start:i])
			}
			start = i + 1
		}
	}
	slr = append(slr, s[start:])
	return slr
}

/*
	s[0:5]
*/
