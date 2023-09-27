package main

import (
	"fmt"
	"os"
)

func main() {
	error := false
	if len(os.Args) < 3 {
		usage()
	}

	convertInt := BasicAtoi(os.Args[2])
	if convertInt <= 0 || convertInt > 2000 {
		usage()
	}

	if os.Args[1] != "-c" {
		fmt.Printf("Usage: ztail [-c N] <file1> <file2> ...\n")
		usage()
	}

	for i := 3; i < len(os.Args); i++ {
		if i > 3 {
			fmt.Printf("\n")
		}
		fmt.Printf("==> %s <==\n", os.Args[i])
		data, err := ReadFile(os.Args[i])
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", os.Args[i], err)
			error := true
			continue
		}
		if len(data) > convertInt {
			data = data[len(data)-convertInt:]
		}
		fmt.Printf("%s", data)
	}
	if error {
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage: ztail [-c N] <file1> <file2> ...\n")
	os.Exit(1)
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

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
