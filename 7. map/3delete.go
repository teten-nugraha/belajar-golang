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

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "januari")
	fmt.Println(len(chicken))
	fmt.Println(chicken)

}
