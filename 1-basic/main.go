package main

import (
	"fmt"

	"github.com/wnfrx/unit-testing-sharing-session/1-basic/utils"
)

func main() {
	x := utils.Sum(3, 5)
	fmt.Printf("Sum of 3 + 5 = %d\n", x)

	y := utils.Max(6, 10)
	fmt.Printf("Max of 6 & 10 = %d\n", y)

	z := utils.Max(10, 0)
	fmt.Printf("Max of 10 & 0 = %d\n", z)
}
