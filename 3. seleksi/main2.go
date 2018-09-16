package main

import (
	"fmt"
)

func main() {
	var point = 8840.5

	if percent := point / 100; percent >= 100 {
		fmt.Println("A")
	} else if percent >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
