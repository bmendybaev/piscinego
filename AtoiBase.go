package piscine

// AtoiBase преобразует строку s в число в указанной системе счисления base
func AtoiBase(s string, base string) int {
	// Проверяем валидность базы
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	// Обрабатываем строку s
	for _, r := range s {
		index := indexInBase(r, base) // Находим индекс символа в базе
		if index == -1 { // Если символ не найден, возвращаем 0
			return 0
		}
		result = result*baseLen + index
	}

	return result
}

// Функция для поиска индекса символа в строке базы
func indexInBase(char rune, base string) int {
	for i, r := range base {
		if r == char {
			return i
		}
	}
	return -1 // Если символ не найден, возвращаем -1
}

// Проверка валидности базы
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i, r1 := range base {
		if r1 == '+' || r1 == '-' {
			return false
		}
		// Проверяем, нет ли дубликатов символов в базе
		for j, r2 := range base {
			if i != j && r1 == r2 {
				return false
			}
		}
	}
	return true
}
