package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	x := '0'

	for x <= '9' {
		y := '0'
		for y <= '9' {
			z := '0'
			for z <= '9' {
				if x != y && y != z && x != z && x < y && y < z {
					z01.PrintRune(rune(x))
					z01.PrintRune(rune(y))
					z01.PrintRune(rune(z))
					if x != '7' || y != '8' || z != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('\n')
					}
				}
				z++
			}
			y++
		}
		x++
	}
}
