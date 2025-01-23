package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b  // Результат деления
	*mod = a % b  // Остаток от деления
}
