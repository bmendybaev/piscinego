package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		printIt("Error\n")
	} else {
		iPos := 0
		iCount := 0
		iCheck := 0
		iErrCount := 0
		for idx, v := range arguments {
			if idx == 0 && v == "-c" {
				iPos = Atoi(arguments[1])
				iCheck = 2
			}
			if idx >= iCheck {
				str := getData(v)
				strErr := "open "
				strErr += v
				strErr += ": no such file or directory\n"
				if str != strErr {
					str = getSubStr(str, iPos)
					sHead := ""
					if iCount > 0 || iErrCount > 0 {
						sHead = "\n"
					}
					sHead += "==> "
					sHead += v
					sHead += " <==\n"
					str = sHead + str
					printIt(str)
					iCount++
				} else {
					printIt(str)
					iErrCount++
				}
			}
		}
		if iErrCount > 0 {
			os.Exit(1)
		}
	}
}

func getSubStr(s string, i int) string {
	output := []byte(s)
	sOut := ""
	iStart := 0
	if i < len(s) {
		iStart = len(s) - i
	}
	for j := iStart; j < len(s); j++ {
		sOut += string(output[j])
	}
	return sOut
}

func getData(s string) string {
	data, err := os.ReadFile(s)
	if err != nil {
		return (err.Error() + "\n")
	}
	return string(data)
}

func printIt(s string) {
	fmt.Printf("%v", s)
}

func Atoi(s string) int {
	var iReturn int
	output := []byte(s)
	var j int
	iMinus := 0
	for i := len(output) - 1; i >= 0; i-- {
		iCheckL := CheckExpression2(int(output[i]) - '0')
		if iCheckL == 0 {
			return 0 // invalid character
		} else if iCheckL == -1 {
			if iMinus == 0 && i == 0 {
				iReturn *= -1 // handle negative sign
			} else {
				return 0
			}
			iMinus++
		} else if iCheckL == -2 {
			if iMinus == 0 && i == 0 {
				iMinus++
			} else {
				return 0
			}
		} else {
			decem := TenStep3(j)
			if j == 0 {
				iReturn += (int(output[i]) - '0')
			} else {
				iReturn += (int(output[i]) - '0') * decem
			}
		}
		j++
	}
	return iReturn
}

func CheckExpression2(iAscii int) int {
	if iAscii >= 0 && iAscii <= 9 {
		return 1
	} else if iAscii == -3 { // handling '-'
		return -1
	} else if iAscii == -5 { // handling '+'
		return -2
	}
	return 0
}

// ddsdfs
func TenStep3(j int) int {
	if j == 0 {
		return 1
	}
	result := 1
	for i := 0; i < j; i++ {
		result *= 10
	}
	return result
}
