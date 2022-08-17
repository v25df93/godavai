package main

import (
	"fmt"
)

func main() {
	i := "davai govno"
	fmt.Scan(&i)
	switch i {
	case "if":
		fmt.Println("if > else")
	case "else":
		fmt.Println("net epta ina4e")
	default:
		fmt.Println("GEIII")
	}
}
