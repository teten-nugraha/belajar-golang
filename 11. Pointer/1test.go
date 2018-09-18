package main

import (
	"fmt"
)

/*
	tanda * = mengambil nilai asli dari pointernya
	tanga & = mengambil nilai pointernya
*/
func main() {
	var numberA = 4
	var numberB *int = &numberA

	fmt.Println("NumberA (value) ", numberA)
	fmt.Println("NumberA (ref) ", &numberA)

	fmt.Println("NumberB (val) ", *numberB)
	fmt.Println("NumberB (ref) ", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("NumberA (value) ", numberA)
	fmt.Println("NumberA (ref) ", &numberA)

	fmt.Println("NumberB (val) ", *numberB)
	fmt.Println("NumberB (ref) ", numberB)

}
