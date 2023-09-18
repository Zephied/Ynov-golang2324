package piscine

func SplitWhiteSpaces(s string) []string {
	tab := make([]string, 0)
	i := 0
	for range s {
		j := i
		for s[j] != ' ' && s[i] != '	' && s[i] != '\n' && j != len(s)-1 {
			j++
		}
		tab = append(tab, s[i:j])
		i = j
		if s[i] == ' ' || s[i] == '	' || s[i] == '\n' {
			for s[i] == ' ' || s[i] == '	' || s[i] == '\n' {
				i++
			}
			tab = append(tab, ", ")
		}
	}
	return tab
}

/*
	s[0:5]
*/
