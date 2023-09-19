package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	r := ""
	i := 0
	for range a {
		r += string(a[i])
		if i != len(a)-1 {
			r += "\n"
		}
		i++
	}
	i = 0
	for range r {
		z01.PrintRune(rune(r[i]))
		i++
	}
}
