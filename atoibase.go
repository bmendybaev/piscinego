import "github.com/01-edu/z01"

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digits = append([]rune{rune(n%10) + '0'}, digits...)
		n /= 10
	}

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}
