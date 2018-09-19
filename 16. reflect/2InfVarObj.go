package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectVal = reflect.ValueOf(s)

	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	var reflectType = reflectVal.Type()

	for i := 0; i < reflectVal.NumField(); i++ {
		fmt.Println("nama		: ", reflectType.Field(i).Name)
		fmt.Println("Tipe Data	: ", reflectType.Field(i).Type)
		fmt.Println("Nilai		: ", reflectVal.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "Teten Nugraha", Grade: 100}
	// s1.getPropertyInfo()
	fmt.Println("name ", s1.Name)

	var reflectVal = reflect.ValueOf(s1)
	var method = reflectVal.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Nugraha"),
	})
	fmt.Println("nama : ", s1.Name)
}
