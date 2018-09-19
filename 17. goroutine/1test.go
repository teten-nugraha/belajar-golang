package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "Teten")
	print(5, "Apa kabar")

	var input string
	fmt.Scanln(&input)
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
