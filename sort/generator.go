package sort

import (
	"math/rand"
	"time"
)

// GenerateSlice generates a slice of random ints of a given size
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
