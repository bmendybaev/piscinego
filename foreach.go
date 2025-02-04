package piscine

import "fmt"

// ForEach применяет переданную функцию f ко всем элементам среза a.
func ForEach(f func(int), a []int) {
	// Проходим по каждому элементу среза
	for _, value := range a {
		// Вызываем функцию f для текущего элемента
		f(value)
	}
}

// PrintNbr выводит число в stdout (имитация PrintNbr из piscine)
func PrintNbr(n int) {
	fmt.Print(n)
}
