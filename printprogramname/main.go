package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := 0
	prog := os.Args[0]
	for range prog {
		z01.PrintRune(rune(prog[i]))
		i++
	}
}
