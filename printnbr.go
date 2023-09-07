package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	tab := make([]int, 0)
	i := 0
	if n+'0' < '0' {
		z01.PrintRune('-')
		n *= -1
	}
	for n >= 10 {
		tab = append(tab, n%10+'0')
		n = n / 10
		i++
	}
	i--
	z01.PrintRune(rune(n + '0'))
	for i >= 0 {
		z01.PrintRune(rune(tab[i]))
		i--
	}
}
