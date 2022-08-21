package main

import (
	"fmt"
)

func joinTwoStrings(prefix string, suffix string) string {
	fmt.Println(prefix, suffix)
	return prefix + suffix
}

func main() {
	joinTwoStrings("vasil'ev", "artem")
}
