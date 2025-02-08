package piscine

import (
	"github.com/01-edu/z01"
	"fmt"
)

func DealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := len(deck) / players

	for i := 0; i < players; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			index := i*cardsPerPlayer + j
			if j == cardsPerPlayer-1 {
				printNumber(deck[index])
			} else {
				printNumber(deck[index])
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	num := []rune{}
	for n > 0 {
		digit := n % 10
		num = append([]rune{rune('0' + digit)}, num...)
		n /= 10
	}

	for _, r := range num {
		z01.PrintRune(r)
	}
}
