package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Store the original values in temporary variables
	tempA := ***a
	tempB := *b
	tempC := *******c
	tempD := ****d

	// Move the values around as per the instructions
	***a = tempB   // b into a
	*b = tempD     // d into b
	*******c = tempA // a into c
	****d = tempC   // c into d
}
