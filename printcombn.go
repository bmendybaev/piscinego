package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// dopustim predel
	if n <= 0 || n >= 10 {
		return
	}

	// massiv
	combination := make([]int, n)
	for i := 0; i < n; i++ {
		combination[i] = i
	}

	for {
		// na pechat
		for _, digit := range combination {
			z01.PrintRune('0' + rune(digit))
		}

		// ogranichitel po kombinatsiayam
		if combination[0] == 10-n {
			break
		}

		// zapyat and probel
		z01.PrintRune(',')
		z01.PrintRune(' ')

		// drugoi uroven
		i := n - 1
		for i >= 0 && combination[i] == 9-(n-1-i) {
			i--
		}
		combination[i]++
		for j := i + 1; j < n; j++ {
			combination[j] = combination[j-1] + 1
		}
	}

	// konec stroki
	z01.PrintRune('\n')
}
