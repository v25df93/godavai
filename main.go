package main

import (
	"errors"
	"fmt"
	"log"
)

func enterTheClub(age int) (string, error) {
	if age >= 18 && age <= 45 {
		return "Входи.", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понравится?", nil
	} else if age > 65 {
		return "Это уже слишком", errors.New("you are too old")
	}

	return "Нельзя", errors.New("you are too young")

}

func main() {
	message, err := enterTheClub(12)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}

	fmt.Println(message)
}
