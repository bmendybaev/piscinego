package piscine

import "fmt"

// ForEach применяет функцию f к каждому элементу среза asdfsdfsdfsd
func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}

// PrintNbr выводит число без новой строки dsfsdfsdfsdfsdfsdfs
func PrintNbr(n int) {
	fmt.Print(n)
}
