package piscine

func Compare(a, b string) int {
	ai, bi := 0, 0
	alen, blen := len(a), len(b)

	// Итерируемся по строкам посимвольно
	for ai < alen && bi < blen {
		if a[ai] < b[bi] {
			return -1
		} else if a[ai] > b[bi] {
			return 1
		}
		ai++
		bi++
	}

	// Если строки равны по символам, но разной длины
	if alen < blen {
		return -1
	} else if alen > blen {
		return 1
	}

	// Если строки идентичны
	return 0
}
