package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	players := [4][3]int{
		{deck[0], deck[1], deck[2]},
		{deck[3], deck[4], deck[5]},
		{deck[6], deck[7], deck[8]},
		{deck[9], deck[10], deck[11]},
	}

	digitRunes := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < 3; j++ {
			n := players[i][j]
			if n == 0 {
				z01.PrintRune('0')
			} else {
				d := 1
				for n/d >= 10 {
					d *= 10
				}
				for d > 0 {
					z01.PrintRune(digitRunes[n/d])
					n %= d
					d /= 10
				}
			}
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
