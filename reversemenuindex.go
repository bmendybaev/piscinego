package piscine

func ReverseMenuIndex(menu []string) []string {
    // Create a new slice with the same length as the input slice
    reversedMenu := make([]string, len(menu))

    // Iterate over the original menu and assign items in reverse order
    for i := 0; i < len(menu); i++ {
        reversedMenu[i] = menu[len(menu)-1-i]
    }

    return reversedMenu
}
