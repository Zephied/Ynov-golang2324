package main

import (
	"fmt"
	"os"
)

/*func main() {
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
				fmt.Printf("open %s: no such file or directory\n", name)
				error = true
			} else {
				j = val
				if val > len(count) {
					j = len(count)
				}
				tail = string(count[len(count)-j:])
				fmt.Printf("%s %s %s\n", start, name, end)
				fmt.Printf("%s", tail)
			}
			i++
		}
		if error {
			os.Exit(1)
		}
	}
}*/

func main() {
	CheckInput()
	//start := "==>"
	//end := "<=="
	//error := false
	//convertInt := BasicAtoi(os.Args[2])
	for i := 3; i < len(os.Args); i++ {
		ReadFile(os.Args[i], os.Args[2])
		/*if convertInt > len(file) {
			convertInt = len(file)
			error = true
		}
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", name)
			error = true
		} else {
			if i > 3 {
				fmt.Printf("\n")
			}
			fmt.Printf("%s %s %s\n", start, name, end)
			fmt.Printf("%s\n", file[len(file)-convertInt:])
		}*/
	}
	/*if error {
		os.Exit(1)
	}*/
}

func CheckInput() {
	convertInt := BasicAtoi(os.Args[2])
	if len(os.Args) < 4 {
		fmt.Printf("Usage: ztail [-c N] <file1> <file2> ...\n")
		os.Exit(1)
	}
	if os.Args[1] != "-c" {
		fmt.Printf("Usage: ztail [-c N] <file1> <file2> ...\n")
		os.Exit(1)
	}
	if convertInt < 0 || convertInt > 2000 {
		fmt.Printf("Usage: ztail [-c N] <file1> <file2> ...\n")
		os.Exit(1)
	}
}

func BasicAtoi(s string) int {
	var result int
	for _, v := range s {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		}
	}
	return result
}

func ReadFile(file string, nbr string) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		fmt.Printf("open %s: no such file or directory\n\n", file)
	} else {
		fileRead, _ := os.ReadFile(file)
		fmt.Printf("%s %s %s\n", "==>", file, "<==")
		fmt.Printf("%s\n", fileRead[len(fileRead)-BasicAtoi(nbr):])
	}
}
