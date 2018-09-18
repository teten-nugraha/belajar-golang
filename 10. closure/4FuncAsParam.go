package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContains0 = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	fmt.Println("filter ada huruf\"o\"\t:", dataContains0)
	fmt.Println("filter jumlah huruf\"5\"\t:", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
