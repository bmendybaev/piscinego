package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Преобразуем из baseFrom в десятичную систему счисления
	decimalValue := baseToDecimal(nbr, baseFrom)

	// Преобразуем из десятичной системы в baseTo
	return decimalToBase(decimalValue, baseTo)
}

// Преобразует число из указанной системы счисления в десятичную
func baseToDecimal(nbr, base string) int {
	baseLen := len(base)
	result := 0

	for _, char := range nbr {
		// Получаем индекс символа в строке base ()
		index := indexOf(char, base)
		result = result*baseLen + index
	}

	return result
}

// Преобразует десятичное число в целевую систему счисления
func decimalToBase(nbr int, base string) string {
	if nbr == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for nbr > 0 {
		result = string(base[nbr%baseLen]) + result
		nbr /= baseLen
	}

	return result
}

// Функция для поиска индекса символа в строке ()
func indexOf(char rune, base string) int {
	for i, c := range base {
		if c == char {
			return i
		}
	}
	return -1 // Если символ не найден (но по условию задачи это исключено)
}
