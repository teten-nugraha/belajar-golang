package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{}
	s1.person = person{"wick", 23}
	s1.grade = 2

	fmt.Println("Name ", s1.person.name)
	fmt.Println("age ", s1.person.age)
	fmt.Println("grade ", s1.grade)
}
