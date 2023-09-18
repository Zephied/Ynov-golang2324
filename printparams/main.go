package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, v := range args {
		for _, c := range v {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
