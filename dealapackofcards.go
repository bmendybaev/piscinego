package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	fmt.Printf("Player 1: ")
	printDigits(deck[0])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[1])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[2])
	z01.PrintRune('\n')

	fmt.Printf("Player 2: ")
	printDigits(deck[3])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[4])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[5])
	z01.PrintRune('\n')

	fmt.Printf("Player 3: ")
	printDigits(deck[6])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[7])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[8])
	z01.PrintRune('\n')

	fmt.Printf("Player 4: ")
	printDigits(deck[9])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[10])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	printDigits(deck[11])
	z01.PrintRune('\n')
}

func printDigits(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []int{}
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}

	for _, d := range digits {
		switch d {
		case 0:
			z01.PrintRune('0')
		case 1:
			z01.PrintRune('1')
		case 2:
			z01.PrintRune('2')
		case 3:
			z01.PrintRune('3')
		case 4:
			z01.PrintRune('4')
		case 5:
			z01.PrintRune('5')
		case 6:
			z01.PrintRune('6')
		case 7:
			z01.PrintRune('7')
		case 8:
			z01.PrintRune('8')
		case 9:
			z01.PrintRune('9')
		}
	}
}
