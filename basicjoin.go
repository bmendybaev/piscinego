package piscine

func BasicJoin(elems []string) string {
	result := ""

	// Iterate through the slice of strings and concatenate each element
	for i := 0; i < len(elems); i++ {
		result += elems[i]
	}

	return result
}
