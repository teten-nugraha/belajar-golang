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

	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}
}
