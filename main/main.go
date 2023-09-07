package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Serzerzeralut!", "derzerzeralutt!"))
	fmt.Println(piscine.Compare("Serzerzeralut!", "Serzerzerayut"))
}
