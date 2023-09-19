package piscine

func SplitWhiteSpaces(s string) []string {
	tab := make([]string, 0)
	i := 0
	for range s {
		j := i
		if s[j] == 0 && s[i] == ' ' || s[j] == 0 && s[i] == '	' || s[j] == 0 && s[i] == '\n' {
			i++
			j++
		}
		for s[j] != ' ' && s[j] != '	' && s[j] != '\n' && j != len(s)-1 {
			j++
		}
		tab = append(tab, s[i:j])
		i = j
		if s[i] == ' ' || s[i] == '	' || s[i] == '\n' {
			for s[i] == ' ' || s[i] == '	' || s[i] == '\n' || i == len(s)-1 {
				i++
			}
		}
	}
	return tab
}

/*
	s[0:5]
*/
