package numbers

import (
	"github.com/Neels-v-Wyk/intro-to-algorithms-4th-ed-algos/go/generators"
)

func Sort() []int {
	// TODO:
	//     don't generate a random number sequence each time, just generate it once in main.go
	random_sequence := generators.GenerateNumbers(3, true)

	return random_sequence
}

