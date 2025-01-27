package piscine

func Capitalize(s string) string {
	result := []byte(s)
	isStartOfWord := true

	for i := 0; i < len(result); i++ {
		if (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= 'A' && result[i] <= 'Z') || (result[i] >= '0' && result[i] <= '9') {
			if isStartOfWord && result[i] >= 'a' && result[i] <= 'z' {
				result[i] -= 32 // Convert lowercase to uppercase
			} else if !isStartOfWord && result[i] >= 'A' && result[i] <= 'Z' {
				result[i] += 32 // Convert uppercase to lowercase
			}
			isStartOfWord = false
		} else {
			isStartOfWord = true // Non-alphanumeric character resets word start
		}
	}

	return string(result)
}
