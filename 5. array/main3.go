package main

import (
	"fmt"
)

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Printf("Elemen  %s \n", fruit)
	}
}
