package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 {
		z01.PrintRune('\n')
		return
	}
	if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}
	for i := '0'; i <= '9'; i++ {
		recursivePrintCombN(n-1, i, i)
	}
}

func recursivePrintCombN(n int, firstRune rune, currentRune rune) {
	if n == 0 {
		z01.PrintRune(currentRune)
		if currentRune != '9' {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
		return
	}
	for i := firstRune + 1; i <= '9'; i++ {
		recursivePrintCombN(n-1, firstRune, i)
	}
}
