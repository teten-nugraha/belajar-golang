package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hallo ", s.name)
}

func (s student) getNameAt(i int) string {
	/*
		strings.Split(s.name, " ") = memecah s.name berdasarkan space menjadi array
		[i-1] adlah index untuk return value
	*/
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{"Teten Nugraha", 28}
	s1.sayHello()

	var name = s1.getNameAt(1)
	fmt.Println("nama panggilan : ", name)
}
