package main

import (
	"fmt"
)

func main() {
	var point = 5

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("Bad")
	}
}
