package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1

	// Iterate through each character in the string
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && result == 0 { // Check for negative sign
			sign = -1
		} else if s[i] >= '0' && s[i] <= '9' { // Check for digits
			result = result*10 + int(s[i]-'0') // Convert character to digit and add to result
		}
	}

	// Return the result with the appropriate sign
	return result * sign
}
