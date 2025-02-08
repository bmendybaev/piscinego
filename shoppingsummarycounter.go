package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	itemCount := make(map[string]int)

	item := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			item += string(str[i])
		} else if item != "" {
			itemCount[item]++
			item = ""
		}
	}
	if item != "" {
		itemCount[item]++
	}

	return itemCount
}
