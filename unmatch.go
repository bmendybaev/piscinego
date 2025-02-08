package piscine

func Unmatch(a []int) int {
	countMap := make(map[int]int)

	// Count occurrences of each number
	for _, num := range a {
		countMap[num]++
	}

	// Find the first number with an odd count in the original order
	for _, num := range a {
		if countMap[num]%2 != 0 {
			return num
		}
	}

	// If all numbers have pairs
	return -1
}
