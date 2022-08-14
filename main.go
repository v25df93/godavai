package main

import "fmt"

func main() {
	call := func(name string) {
		fmt.Println("Calling", name)
	}
	call("Meet")
	person := "Chris"
	call(person)
}
