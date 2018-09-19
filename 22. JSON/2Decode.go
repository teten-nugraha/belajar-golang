package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `[
		{"Name":"Teten Nugraha","Age":27},
		{"Name":"Yasna Almahyra","Age":1}
	]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User 1: ", data[0].FullName)
	fmt.Println("User 2: ", data[1].FullName)
}
