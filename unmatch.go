package piscine

func Unmatch(a []int) int {
	countMap := make(map[int]int)

	// Count occurrences of each number
	for _, num := range a {
		countMap[num]++
	}

	// Find the number with an odd count
	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}

	// If all numbers have pairs
	return -1
}
