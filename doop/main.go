package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	num1, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	if err1 != nil || err2 != nil {
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
			fmt.Println("No division by 0")
			return
		}
		result = num1 / num2
	case "%":
		if num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = num1 % num2
	default:
		return
	}

	if overflow {
		return
	}

	fmt.Println(result)
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
