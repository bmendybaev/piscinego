package piscine

import (
	"fmt"
)

// EightQueens запускает решение задачи восьми ферзей
func EightQueens() {
	solve("", 0)
}

// solve рекурсивно строит все возможные решения
func solve(current string, row int) {
	// Если 8 ферзей расставлены, выводим решение
	if row == 8 {
		fmt.Println(current)
		return
	}

	// Проверяем все возможные столбцы для текущей строки
	for col := 1; col <= 8; col++ {
		if isValid(current, row, col) {
			// Добавляем текущий столбец в строку решения
			solve(current+fmt.Sprint(col), row+1)
		}
	}
}

// isValid проверяет, можно ли поставить ферзя в (row, col)
func isValid(current string, row, col int) bool {
	for r := 0; r < row; r++ {
		// Получаем столбец ферзя на предыдущих строках
		placedCol := int(current[r] - '0')

		// Проверяем:
		// - Ферзь не находится в том же столбце
		// - Ферзь не находится на диагонали
		if placedCol == col || abs(placedCol-col) == abs(r-row) {
			return false
		}
	}
	return true
}

// abs возвращает модуль числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
