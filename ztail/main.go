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
	if os.Args[1] == "-c" && os.Args[2] >= "0" && os.Args[2] <= "9" {
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
				fmt.Printf("err")
			}
			j = val
			tail = string(count[len(count)-j:])
			fmt.Printf("%s %s %s \n", start, name, end)
			fmt.Printf("%s", tail)
			i++
		}
	}
}
