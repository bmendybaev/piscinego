package main

import (
	"fmt"

	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrimeCheckcd, a)
	fmt.Println(result)
}
