package main

import (
	"fmt"
	"os"
)

func main() {
	i := 3
	var name string
	start := "==>"
	end := "<=="
	var j int
	error := false
	if os.Args[1] == "-c" && os.Args[2] >= "0" && os.Args[2] <= "2345" {
		val := 0
		for _, va := range os.Args[2] {
			if va >= '0' && va <= '9' {
				val = val*10 + int(va-'0')
			}
		}
		for i < len(os.Args) {
			tail := ""
			count, err := os.ReadFile(os.Args[i])
			name = os.Args[i]
			if err != nil {
				fmt.Printf("open %s: no such file or directory\n\n", name)
				error = true
			} else {
				j = val
				if val > len(count) {
					j = len(count)
				}
				tail = string(count[len(count)-j:])
				fmt.Printf("%s %s %s \n", start, name, end)
				fmt.Printf("%s", tail)
			}
			i++
		}
		if error {
			os.Exit(1)
		}
	}
}
