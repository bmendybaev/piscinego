package main

import (
	"os"

	"github.com/01-edu/z01"
)

type boolean bool

const (
	yes     boolean = true
	no      boolean = false
	EvenMsg         = "I have an even number of arguments"
	OddMsg          = "I have an odd number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if nbr%2 == 0 {
		return yes
	}
	return no
}

func main() {
	lengthOfArg := len(os.Args[1:])

	if isEven(lengthOfArg) == yes {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
