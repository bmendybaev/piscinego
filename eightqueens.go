package piscine

import "fmt"

// EightQueens prints all solutions to the eight queens puzzle
func EightQueens() {
	solveQueens("", 0)
}

// solveQueens recursively tries to solve the puzzle
func solveQueens(current string, row int) {
	// Base case: If all 8 queens are placed, print the solution
	if row == 8 {
		fmt.Println(current)
		return
	}

	// Try placing a queen in each column for the current row
	for col := 1; col <= 8; col++ {
		if isValidPosition(current, row, col) {
			// Add the current column to the solution and recurse
			solveQueens(current+fmt.Sprint(col), row+1)
		}
	}
}

// isValidPosition checks if a queen can be placed at (row, col)
func isValidPosition(current string, row, col int) bool {
	for r := 0; r < row; r++ {
		// Column of the previously placed queen
		placedCol := int(current[r] - '0')

		// Check for conflicts:
		// - Same column
		// - Same diagonal
		if placedCol == col || abs(placedCol-col) == abs(r-row) {
			return false
		}
	}
	return true
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
