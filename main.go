package main

import "fmt"

func swap(a int, b int) {
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}

func main() {
	a := 10
	b := 4
	swap(a, b)
}
