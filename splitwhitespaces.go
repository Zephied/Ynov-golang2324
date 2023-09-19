package piscine

/*
	func SplitWhiteSpaces(s string) []string {
		tab := make([]string, 0)
		i := 0
		for range s {
			j := i
			if s[i] != ' ' && s[i] != '	' && s[i] != '\n' {
				for s[j] != ' ' && s[j] != '	' && s[j] != '\n' && j < len(s)-1 {
					j++
				}
				if i != j && s[i:j] != "" {
					tab = append(tab, s[i:j])
				}
				i = j
				for s[i] == ' ' || s[i] == '	' || s[i] == '\n' || i == len(s) {
					i++
				}
			} else {
				i++
			}
		}
		return tab
	}

s[0:5]

	func SplitWhiteSpaces(s string) []string {
		start := 0
		slr := []string{}
		for i, v := range s {
			if v == ' ' {
				if i != start && s[i:start] != "" {
					slr = append(slr, s[start:i])
				}
				start = i + 1
			}
		}
		slr = append(slr, s[start:])
		return slr
	}
*/
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
	if s[start:] != "" {
		slr = append(slr, s[start:])
	}
	return slr
}
