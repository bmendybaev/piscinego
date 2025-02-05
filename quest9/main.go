package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{5, 4, 3, 2, 1, 0}
	a4 := []int{951689, 847168, 788568, -141653, -228980, -416475, -587483, -847340}

	result1 := piscine.IsSorted(piscine.IsSortedBy10, a1)
	result2 := piscine.IsSorted(piscine.IsSortedBy10, a2)
	result3 := piscine.IsSorted(piscine.IsSortedBy10, a3)
	result4 := piscine.IsSorted(piscine.IsSortedBy10, a4)

	fmt.Println(result1) // true
	fmt.Println(result2) // false
	fmt.Println(result3) // true
	fmt.Println(result4) // true или false в зависимости от правильности логики
}
