package generators

import (
	"math/rand"
	"time"
)

// Generates a list of randomly organized numbers
// by default it doesn't create duplicate numbers in this sequence
func GenerateNumbers(size int, duplicates_allowed bool) []int {

	empty_array := make([]int, size)
    rand.Seed(time.Now().UnixNano())

	if duplicates_allowed {
		for i := 1; i < size; i++ {
			empty_array[i] = rand.Intn(size)
		}
	}

	return empty_array
}
