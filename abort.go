package piscine

func Abort(a, b, c, d, e int) int {
    // Manually find the median using comparisons
    if (a > b) { a, b = b, a }
    if (a > c) { a, c = c, a }
    if (a > d) { a, d = d, a }
    if (a > e) { a, e = e, a }
    if (b > c) { b, c = c, b }
    if (b > d) { b, d = d, b }
    if (b > e) { b, e = e, b }
    if (c > d) { c, d = d, c }
    if (c > e) { c, e = e, c }
    if (d > e) { d, e = e, d }

    // The median is now the middle value, which is c
    return c
}
