package piscine

import (
	"strings"
)

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	items := strings.Split(str, " ")

	for _, item := range items {
		summary[item]++
	}

	return summary
}
