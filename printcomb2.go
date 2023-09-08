package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	x := '0'
	for x <= '9' {
		y := '0'
		for y <= '9' {
			z := '0'
			for z <= '9' {
				i := '0'
				for i <= '9' {
					if y < i || z > '0' {
						z01.PrintRune(rune(x))
						z01.PrintRune(rune(y))
						z01.PrintRune(rune(' '))
						z01.PrintRune(rune(z))
						z01.PrintRune(rune(i))
						if x != '9' || y != '8' || z != '9' || i != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune('\n')
						}
					}
					i++
				}
				z++
			}
			y++
		}
		x++
	}
}
