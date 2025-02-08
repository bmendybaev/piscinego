
package piscine
func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	s := ""
	c := 3
	for i := 2; i < len(str); i++ {
		if c == 3 {
			s += string(str[i])
			c = 0
		}
		c++
	}
	s += "\n"
	return s
}
