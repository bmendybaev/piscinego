package piscine

import (
	"strings"
)

func ShoppingSummaryCounter(str string) map[string]int {
	// Create a map to store the item counts
	itemCount := make(map[string]int)

	// Split the string by spaces to get individual items
	items := strings.Split(str, " ")

	// Loop through each item and count occurrences
	for _, item := range items {
		itemCount[item]++
	}

	return itemCount
}
