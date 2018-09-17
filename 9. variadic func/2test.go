package main

import (
	"fmt"
	"strings"
)

func main() {
	yourHobbies("code", "sleep", "eat")
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n ", name)
	fmt.Printf("My Hobbies area : %s\n", hobbiesAsString)
}
