package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	sortint := make([]int, 0)
	varstack := make([]int, 0)
	i := 1
	j := 0
	y := 0
	if n != 0 {
		for n > 0 {
			sortint = append(sortint, n%10+'0')
			n = n / 10
		}
		for i < len(sortint) {
			if sortint[i] < sortint[i-1] {
				varstack = append(varstack, sortint[i-1])
				sortint[i-1] = sortint[i]
				sortint[i] = varstack[y]
				y++
				if i != 1 {
					i--
				}
			} else {
				i++
			}
		}
		for j < i {
			z01.PrintRune(rune(sortint[j]))
			j++
		}
	} else {
		z01.PrintRune(rune(n + '0'))
	}
}
