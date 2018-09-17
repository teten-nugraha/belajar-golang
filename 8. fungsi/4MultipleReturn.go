package main

import (
	"fmt"
	"math"
)

/*
(d float64) untuk parameter inputan
(float64, float64) tipe data untuk kembalian
*/
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	var circumrefence = math.Pi * d

	// kembalikan 2 nilai
	return area, circumrefence
}

func main() {
	var diameter float64 = 15
	var area, circumreference = calculate(diameter)

	fmt.Println("luas lingkaran \t ", area)
	fmt.Println("keliling lingkaran ", circumreference)
}
