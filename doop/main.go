package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	num1, ok1 := parseInt(os.Args[1])
	operator := os.Args[2]
	num2, ok2 := parseInt(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	var result int64
	var overflow bool

	switch operator {
	case "+":
		result, overflow = safeAdd(num1, num2)
	case "-":
		result, overflow = safeSub(num1, num2)
	case "*":
		result, overflow = safeMul(num1, num2)
	case "/":
		if num2 == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		result = num1 / num2
	case "%":
		if num2 == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		result = num1 % num2
	default:
		return
	}

	if overflow {
		return
	}

	writeInt(result)
}

func parseInt(s string) (int64, bool) {
	var num int64
	var sign int64 = 1
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		}
		if c < '0' || c > '9' {
			return 0, false
		}
		num = num*10 + int64(c-'0')
	}
	return num * sign, true
}

func writeInt(n int64) {
	var buf [20]byte
	i := len(buf)
	sign := false
	if n < 0 {
		sign = true
		n = -n
	}

	for {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
		if n == 0 {
			break
		}
	}

	if sign {
		i--
		buf[i] = '-'
	}

	os.Stdout.Write(buf[i:])
	os.Stdout.Write([]byte{'
'})
}

func safeAdd(a, b int64) (int64, bool) {
	res := a + b
	if (b > 0 && res < a) || (b < 0 && res > a) {
		return 0, true
	}
	return res, false
}

func safeSub(a, b int64) (int64, bool) {
	res := a - b
	if (b > 0 && res > a) || (b < 0 && res < a) {
		return 0, true
	}
	return res, false
}

func safeMul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}
	res := a * b
	if res/b != a {
		return 0, true
	}
	return res, false
}
