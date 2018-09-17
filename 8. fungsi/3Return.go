package main

import (
	"fmt"
)

func main() {
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Println("Pembagi tidak boleh 0")
		return
	}

	var res = m / n
	fmt.Printf("Nilai pembagi adalah %d \n", res)
}
