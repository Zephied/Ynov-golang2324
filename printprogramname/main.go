package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := 2
	prog := os.Args[0]
	for i < len(prog) {
		z01.PrintRune(rune(prog[i]))
		i++
	}
	z01.PrintRune('\n')
}
