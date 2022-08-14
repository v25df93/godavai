package main

import (
	"fmt"
)

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "Входи.", true
	} else {
		return "Тебе нет 18.", false
	}
}

func main() {
	message, entered := enterTheClub(18)
	fmt.Println(message)
	fmt.Println(entered)

}
