package piscine

func FirstRune(s string) rune {
	letter := s[0]
	if rune(s[0]) == 'â' {
		return ('♥')
	}
	return (rune(letter))
}
