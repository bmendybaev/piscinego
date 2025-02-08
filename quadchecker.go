package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func quadA(w, h int) string {
	var result strings.Builder
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (x == 0 && y == 0) || (x == w-1 && y == 0) || (x == 0 && y == h-1) || (x == w-1 && y == h-1) {
				result.WriteRune('o')
			} else if y == 0 || y == h-1 {
				result.WriteRune('-')
			} else if x == 0 || x == w-1 {
				result.WriteRune('|')
			} else {
				result.WriteRune(' ')
			}
		}
		result.WriteRune('\n')
	}
	return result.String()
}

func quadC(w, h int) string {
	var result strings.Builder
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y == 0 && x == 0 {
				result.WriteRune('A')
			} else if y == 0 && x == w-1 {
				result.WriteRune('A')
			} else if y == h-1 && x == 0 {
				result.WriteRune('C')
			} else if y == h-1 && x == w-1 {
				result.WriteRune('C')
			} else if y == 0 || y == h-1 {
				result.WriteRune('B')
			} else if x == 0 || x == w-1 {
				result.WriteRune('B')
			} else {
				result.WriteRune(' ')
			}
		}
		result.WriteRune('\n')
	}
	return result.String()
}

func quadD(w, h int) string {
	return quadC(w, h)
}

func quadE(w, h int) string {
	return quadC(w, h)
}

var quads = map[string]func(int, int) string{
	"quadA": quadA,
	"quadC": quadC,
	"quadD": quadD,
	"quadE": quadE,
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	inputStr := string(input)
	matches := []string{}

	for name, quadFunc := range quads {
		for h := 1; h <= len(inputStr); h++ {
			for w := 1; w <= len(inputStr); w++ {
				if quadFunc(w, h) == inputStr {
					match := fmt.Sprintf("[%s] [%d] [%d]", name, w, h)
					matches = append(matches, match)
				}
			}
		}
	}

	sort.Strings(matches)

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
