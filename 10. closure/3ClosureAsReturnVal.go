package main

import (
	"fmt"
)

func main() {
	var max = 3
	var numbers = []int{2, 5, 2, 2, 1, 4, 6, 67, 4, 4, 2, 1}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Println("find \t: %d \n\n:", max)
	fmt.Println("found \t:", howMany)
	fmt.Println("value \t", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e >= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
