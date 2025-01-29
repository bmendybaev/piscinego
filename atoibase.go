package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for i := 0; i < len(s); i++ {
		value := indexInBase(s[i], base)
		if value == -1 {
			return 0
		}
		result = result*baseLen + value
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func indexInBase(c byte, base string) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}
