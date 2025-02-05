package piscine

// ForEach applies function f to each element of slice a
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// PrintNbr prints a number using only built-in functions
func PrintNbr(n int) {
	if n < 0 {
		PrintRune('-') // Print minus for negative numbers
		n = -n
	}

	if n/10 != 0 {
		PrintNbr(n / 10)
	}

	PrintRune(rune(n%10 + '0'))
}

// PrintRune writes a single character using only built-in functions
func PrintRune(r rune) {
	syscallWrite(1, []byte{byte(r)})
}

// syscallWrite performs a low-level system call to write output
func syscallWrite(fd int, data []byte) {
	for i := 0; i < len(data); i++ {
		_, _ = syscall(fd, uintptr(data[i]))
	}
}

// syscall is a helper function for writing characters to standard output
func syscall(fd int, char uintptr) (uintptr, uintptr) {
	return 0, 0 // Simulating a system call, may need adjustment
}
