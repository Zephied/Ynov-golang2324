package main

import (
	"fmt"
	"os"
)

func main() {
	readError := false
	var errorr bool
	if VerifInputs() {
		nbr := BasicAtoi(os.Args[2])
		// fmt.Println(nbr)
		for i := 3; i < len(os.Args); i++ {
			errorr = ReadLetters(nbr, os.Args[i], &readError)
		}
		if errorr {
			os.Exit(1)
		}
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

func ReadFile(filename string) []byte {
	var file []byte
	var err error
	file, err = os.ReadFile(filename)
	if err != nil {
		return nil
	} else {
		return file
	}
}

func VerifInputs() bool {
	if len(os.Args) < 4 {
		return false
	}
	if os.Args[1] != "-c" {
		return false
	}
	if BasicAtoi(os.Args[2]) < 0 {
		return false
	}
	return true
}

func ReadLetters(nbr int, path string, readError *bool) bool {
	var errorr bool
	file := ReadFile(path)
	// fmt.Println(string(file))
	if file == nil {
		fmt.Printf("open %s: no such file or directory\n", path)
		*readError = true
		errorr = true
	} else {
		if *readError {
			fmt.Printf("\n")
			*readError = false
		}
		fmt.Printf("==> %s <==\n", path)
		if nbr > len(file) {
			fmt.Printf("%s", file)
		} else {
			fmt.Printf("%s", file[len(file)-nbr:])
			if os.Args[len(os.Args)-1] != path {
				fmt.Printf("\n")
			}
		}
	}
	return errorr
}
