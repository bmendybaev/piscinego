package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	// Test the NRune function with various inputs
	z01.PrintRune(piscine.NRune("Hello!", 3)) // 'l'
	z01.PrintRune(piscine.NRune("Salut!", 2)) // 'a'
	z01.PrintRune(piscine.NRune("Bye!", -1))  // Invalid, prints nothing
	z01.PrintRune(piscine.NRune("Bye!", 5))   // Invalid, prints nothing
	z01.PrintRune(piscine.NRune("Ola!", 4))   // '!'
	z01.PrintRune('\n')                       // Add a newline at the end
}
