package piscine

import "fmt"

// ForEach applies function f to each element of slice a
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// PrintNbr prints an integer without a newline dddddd
func PrintNbr(n int) {
	fmt.Print(n)
}