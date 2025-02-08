package piscine

import (
	"sort"
)

func ShoppingListSort(slice []string) []string {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})

	return slice
}
