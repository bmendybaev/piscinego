package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 { // Обрабатываем некорректные значения
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ { // Итеративно вычисляем факториал
		result *= i
	}

	return result
}
