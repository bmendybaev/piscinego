package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 3, 2, 1, 0}

	result1 := piscine.IsSorted(piscine.F, a1)
	result2 := piscine.IsSorted(piscine.F, a2)
	result3 := piscine.IsSorted(piscine.F, a3)

	fmt.Println(result1) // true
	fmt.Println(result2) // false
	fmt.Println(result3) // false
}
