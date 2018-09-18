package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 87
	s1.grade = 2

	fmt.Println("Name ", s1.name)
	fmt.Println("age ", s1.age)
	fmt.Println("age ", s1.person.age)
	fmt.Println("grade ", s1.grade)
}
