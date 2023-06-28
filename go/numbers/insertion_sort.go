package numbers

import (
	"github.com/Neels-v-Wyk/intro-to-algorithms-4th-ed-algos/go/generators"
	"fmt"
)

func Sort() []int {
	random_sequence := generators.GenerateNumbers(10, true)
	fmt.Println(random_sequence)

	return []int{1, 2, 3, 4}
}

