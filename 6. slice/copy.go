package main

import (
	"fmt"
)

func main() {
	var fruits = []string{"apple"}
	var aFruits = []string{"watermelon", "pinnapke"}

	var copiedElemen = copy(fruits, aFruits)

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(copiedElemen)
}
