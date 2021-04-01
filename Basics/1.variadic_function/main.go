package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println("Total:", sum(1, 2, 3, 4, 5))
}
