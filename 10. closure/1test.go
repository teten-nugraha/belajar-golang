package main

import (
	"fmt"
)

func main() {

	var getMinMax = func(n []int) (min, max int) {

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max

	}

	var numbers = []int{2, 1, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\n min : %v \nmax : %v\n ", numbers, min, max)
}
