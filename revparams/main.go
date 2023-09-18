package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := 0
	for range os.Args {
		i++
	}
	i--
	for i > 0 {
		args := os.Args[i]
		for _, v := range args {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		i--
	}
}
