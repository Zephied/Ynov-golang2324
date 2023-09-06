package main

import "github.com/01-edu/z01"

func main() {
	a := 'a'
	for a <= 'z' {
		z01.PrintRune(a)
		a++
	}
	z01.PrintRune('\n')
}
