package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, nums := range nums {
		total += nums
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(7, 2, 5, 4))
	fmt.Println(sum())
}
