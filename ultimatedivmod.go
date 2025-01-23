package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b // Вычисляем результат деления
	mod := *a % *b // Вычисляем остаток от деления
	*a = div // Сохраняем результат в a
	*b = mod // Сохраняем остаток в b
}