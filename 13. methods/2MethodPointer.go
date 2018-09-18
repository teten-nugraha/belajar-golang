package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func (s student) changeName1(name string) {
	fmt.Println("on changeName1, name changed to ", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("on changeName2, name changed to ", name)
	s.name = name
}

func main() {
	var s1 = student{"Teten Nugraha", 26}
	fmt.Println("s1 before ", s1.name)

	s1.changeName1("Yasna")
	fmt.Println("s1 after changeName1 ", s1.name)

	s1.changeName2("Abbu")
	fmt.Println("s1 after changeName2 ", s1.name)
}
