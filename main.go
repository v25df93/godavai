package main

import (
	"fmt"
)

func printMessage(message string) {
	fmt.Println(message)
}

func sayhello(name string) string {
	return "привет, " + name + "!!!"
}

func main() {
	message := sayhello("Максим")
	printMessage(message)
}
