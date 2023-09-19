package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if base == "" || StrLen(base) == 1 {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
	if nbr < 0 {
		z01.PrintRune('-')
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
	}
	if nbr < 0 {
		nbr = -nbr
	}
	if nbr >= StrLen(base) {
		PrintNbrBase(nbr/StrLen(base), base)
	}
	z01.PrintRune(rune(base[nbr%StrLen(base)]))
}
