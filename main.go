package main

import (
	f "fmt"
)

func main() {
	f.Println(findmax(1, 2, 5, 7, 788, 535, 321, -2, -01, -55, 2527))
}

func findmax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}

	return max
}
