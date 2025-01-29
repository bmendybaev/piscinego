package piscine

// AtoiBase преобразует строку s в число в указанной системе счисления base
func AtoiBase(s string, base string) int {
	// Проверяем валидность базы
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)

	// Создаём мапу для быстрого поиска индексов символов в базе
	baseMap := make(map[rune]int)
	for i, r := range base {
		baseMap[r] = i
	}

	result := 0
	// Обрабатываем строку s
	for _, r := range s {
		val, exists := baseMap[r]
		if !exists { // Если символ не найден в базе, число недействительно
			return 0
		}
		result = result*baseLen + val
	}

	return result
}

// Проверка валидности базы
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}
