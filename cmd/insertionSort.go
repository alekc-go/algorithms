package main

import (
	"algorithms/sort/insertion"
	"fmt"
)

func main() {
	var input = []int{3, 8, 5, 4, 1, 9, -2}

	fmt.Printf("Insert Sort %v\n", insertion.Sort(input))
	fmt.Printf("Recursive Insert Sort %v\n", insertion.SortRecursive(input, 1, len(input)-1))
}
