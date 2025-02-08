package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	player1, player2, player3, player4 := deck[0], deck[1], deck[2], deck[3]
	player5, player6, player7, player8 := deck[4], deck[5], deck[6], deck[7]
	player9, player10, player11, player12 := deck[8], deck[9], deck[10], deck[11]

	fmt.Printf("Player 1: ")
	z01.PrintRune(rune('0' + player1))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player2))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player3))
	z01.PrintRune('\n')

	fmt.Printf("Player 2: ")
	z01.PrintRune(rune('0' + player4))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player5))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player6))
	z01.PrintRune('\n')

	fmt.Printf("Player 3: ")
	z01.PrintRune(rune('0' + player7))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player8))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player9))
	z01.PrintRune('\n')

	fmt.Printf("Player 4: ")
	z01.PrintRune(rune('0' + player10))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player11))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune(rune('0' + player12))
	z01.PrintRune('\n')
}
