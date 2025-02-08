package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	// Define the menu with corresponding preparation times
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	// Check if the order exists in the menu
	item, exists := menu[order]
	if !exists {
		// Return error code 404 if the item is not found
		return 404
	}

	// Return the preparation time of the item
	return item.preptime
}
