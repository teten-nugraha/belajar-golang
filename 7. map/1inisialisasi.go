package main

import (
	"fmt"
)

func main() {
	var chicken1 = map[string]int{"januari": 50, "februari": 90}

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 90,
	}

	fmt.Println(chicken1["januari"])
	fmt.Println(chicken1["februari"])
	fmt.Println(chicken2["februari"])
	fmt.Println(chicken2["februari"])
}
