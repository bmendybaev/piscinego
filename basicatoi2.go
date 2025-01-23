package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		// Проверяем, является ли символ цифрой
		if char < '0' || char > '9' {
			return 0 // Если символ не цифра, возвращаем 0
		}
		// Вычисляем числовое значение символа и добавляем его к результату
		result = result*10 + int(char-'0')
	}
	return result
}
