package piscine

import "github.com/01-edu/z01"

// EightQueens prints all solutions to the eight queens puzzle
func EightQueens() {
	findSolution("", 0)
}

// findSolution recursively tries to solve the puzzle
func findSolution(current string, row int) {
	if row == 8 { // Base case: All 8 queens are placed
		printSolution(current)
		return
	}

	// Try placing a queen in each column
	for col := 1; col <= 8; col++ {
		if isSafe(current, row, col) {
			findSolution(current+string(rune('0'+col)), row+1)
		}
	}
}

// isSafe checks if placing a queen at (row, col) is valid
func isSafe(current string, row, col int) bool {
	for r := 0; r < row; r++ {
		placedCol := int(current[r] - '0')
		if placedCol == col || abs(placedCol-col) == abs(r-row) {
			return false
		}
	}
	return true
}

// printSolution prints the solution using z01.PrintRune
func printSolution(solution string) {
	for _, r := range solution {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// abs returns the absolute value of a number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
