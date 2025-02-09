package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	s := ""
	str1 := []byte(str)
	sw := ""
	iCount := 0

	for i := 0; i < len(str1); i++ {
		if iCount == 5 {
			iCount = 0
			if len(s) > 0 {
				s += " "
				s += sw
			} else {
				s += sw
			}
			sw = ""
		} else {
			if str1[i] != 32 {
				sw += string(str1[i])
				iCount++
			}
		}
	}

	if len(s) > 0 && len(sw) > 0 {
		s += " "
		s += sw
	} else {
		s += sw
	}

	s += "\n"
	return s
}
