package main

import (
	"fmt"
)

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
}
