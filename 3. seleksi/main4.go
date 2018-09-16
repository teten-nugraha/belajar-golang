package main

import (
	"fmt"
)

func main() {
	var point = 6

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("You need to learn")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You need to learn")
		}
	}
}
