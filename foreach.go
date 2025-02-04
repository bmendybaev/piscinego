package main

import "fmt"

// ForEach применяет переданную функцию f ко всем элементам среза a.
func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value) // Вызываем функцию для каждого элемента
	}
}

// PrintNbr печатает число без пробела или новой строки.
func PrintNbr(n int) {
	fmt.Print(n)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a) // Вызываем ForEach с PrintNbr
}
