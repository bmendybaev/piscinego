package piscine
func ShoppingSummaryCounter(str string) map[string]int {
	m := make(map[string]int)
	out1 := []byte(str)
	var sWord []byte
	var out []string
	for i := 0; i < len(out1); i++ {
		if str[i] != ' ' {
			sWord = append(sWord, out1[i])
		} else {
			out = append(out, string(sWord))
			sWord = make([]byte, 0)
		}
	}
	out = append(out, string(sWord))
	for _, v := range out {
		_, ok := m[v]
		if ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	return m
}
