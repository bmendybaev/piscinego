package main

import (
	"fmt"
	"sort"
)

func ShoppingListSort(slice []string) []string {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})

	return slice
}

func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(ShoppingListSort(slice))
}
