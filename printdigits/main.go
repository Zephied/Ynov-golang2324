package main

import "github.com/01-edu/z01"

func main() {
	a := '0'
	for a <= '9' {
		z01.PrintRune(a)
		a++
	}
	z01.PrintRune('\n')
}
