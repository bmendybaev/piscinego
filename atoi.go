package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	i := 0

	// Проверяем первый символ на знак
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	// Проходим по остальным символам
	for ; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0 // Если символ не цифра, возвращаем 0
		}
		result = result*10 + int(char-'0')
	}

	return result * sign
}
