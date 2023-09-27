package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Parse the --upper flag
	upper := flag.Bool("upper", false, "print result in upper case")
	flag.Parse()

	// Loop through the command-line arguments
	for _, arg := range flag.Args() {
		// Convert the argument to an integer
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Print(" ")
			continue
		}

		// Check if n is a valid position in the alphabet
		if n < 1 || n > 26 {
			fmt.Print(" ")
			continue
		}

		// Calculate the corresponding letter and print it
		letter := string('a' + n - 1)
		if *upper {
			letter = strings.ToUpper(letter)
		}
		fmt.Print(letter)
	}
	fmt.Println()
}
