package main

import "fmt"

// Define the struct point
type point struct {
	x int
	y int
}

// Function to set values in the struct
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{} // Create a pointer to a struct

	setPoint(points) // Modify struct values

	// Print values
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
