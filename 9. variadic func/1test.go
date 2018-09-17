package main

import (
	"fmt"
)

func main() {
	// var avg = calculate(2, 3, 2, 1, 3, 2, 2, 1, 1, 4, 5, 2)
	var numbers = []int{2, 3, 2, 1, 3, 2, 2, 1, 1, 4, 5, 2}
	var avg2 = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f ", avg2)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
