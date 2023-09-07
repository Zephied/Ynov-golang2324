package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	i := 0
	for range s {
		z01.PrintRune(rune(s[i]))
		i++
	}
}
