package main

import (
	f "fmt"
)

func main() {
	f.Println(findmin(1, 2, 5, 7, 788, 535, 321, -2, -01, -55, 2527))
}

func findmin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]

	for _, i := range numbers {
		if i > min {
			min = i
		}
	}

	return min
}
