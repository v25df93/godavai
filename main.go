package main

import (
	"fmt"
)

func printMessage(message string) {
	fmt.Println(message)
}

func sayhello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет!", name, age)
}

func main() {
	var message string
	message = sayhello("Максим", 21)

	printMessage(message)
}
