package main

import (
	"errors"
	"fmt"
)

func prediction(dayofweek string) (string, error) {
	switch dayofweek {
	case "пн":
		return "хорошего начала недели", nil
	case "вт":
		return "хорошего вторника", nil
	case "ср":
		return "хорошего среды", nil
	case "чт":
		return "хорошего четверга", nil
	case "пт":
		return "хорошего пятницы", nil
	case "сб":
		return "хорошего субботы", nil
	case "вс":
		return "хорошего воскресенья", nil
	default:
		return "красава", errors.New("invalid dayofweek")
	}
}

func main() {
	prediction("пн")
	fmt.Println(prediction("asdasdasd"))
}
