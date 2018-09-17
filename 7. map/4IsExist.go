package main

import (
	"fmt"
)

func main() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 90,
		"maret":    90,
		"april":    90,
	}

	var value, isExists = chicken["april"]

	if isExists {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}
