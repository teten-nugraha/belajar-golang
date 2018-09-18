package library

import (
	"fmt"
)

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Teten Nugraha"
	Student.Grade = 2

	fmt.Println("---> library/library.go imported")
}
