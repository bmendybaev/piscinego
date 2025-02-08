package piscine

func StringToIntSlice(str string) []int {
    var intSlice []int
    for _, ch := range str {
        intSlice = append(intSlice, int(ch))
    }
    return intSlice
}
