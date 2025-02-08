package piscine

func Rot14(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			shifted := 'a' + (char-'a'+14)%26
			result += string(shifted)
		} else if char >= 'A' && char <= 'Z' {
			shifted := 'A' + (char-'A'+14)%26
			result += string(shifted)
		} else {
			result += string(char)
		}
	}
	return result
}
