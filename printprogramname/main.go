package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, v := range os.Args[0][2:] + "\n" {
		z01.PrintRune(v)
	}
}
