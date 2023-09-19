package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {

					if i*10+j < k*10+l {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(' ')
						z01.PrintRune(rune(k))
						z01.PrintRune(rune(l))
						if i != 9 || j != 8 || k != 9 || l != 9 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
