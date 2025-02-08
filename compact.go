package piscine

func Compact(ptr *[]string) int {
    slice := *ptr
    compacted := make([]string, 0, len(slice))

    for _, v := range slice {
        if v != "" {
            compacted = append(compacted, v)
        }
    }

    *ptr = compacted
    return len(compacted)
}
