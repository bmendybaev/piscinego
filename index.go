package piscine

func Index(s string, toFind string) int {
	// Если искомая строка пустая, возвращаем 0
	if toFind == "" {
		return 0
	}

	// Перебираем каждый символ строки s
	for i := 0; i < len(s); i++ {
		// Проверяем, хватает ли оставшихся символов для совпадения
		if i+len(toFind) > len(s) {
			return -1
		}

		// Проверяем каждый символ строки toFind
		match := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}

		// Если нашли совпадение, возвращаем индекс
		if match {
			return i
		}
	}

	// Если совпадение не найдено, возвращаем -1
	return -1
}
