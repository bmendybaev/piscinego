package piscine

import (
	"sort"
)

// AdvancedSortWordArr sorts a slice of strings based on the provided comparison function.
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	sort.Slice(a, func(i, j int) bool {
		return f(a[i], a[j]) < 0
	})
}
