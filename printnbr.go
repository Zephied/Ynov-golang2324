package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	tab := make([]int, 0)
	i := 0
	if n < i {
		z01.PrintRune('-')
		for n != 0 {
			tab = append(tab, n%10*(-1))
			n = n / 10
			i++
		}
	} else {
		for n != 0 {
			tab = append(tab, n%10)
			n = n / 10
			i++
		}
	}
	i--
	for i >= 0 {
		z01.PrintRune(rune(tab[i] + '0'))
		i--
	}
}
