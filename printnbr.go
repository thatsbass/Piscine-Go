package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = 223372036854775808
	}

	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}

	digit := '0' + rune(n%10)
	z01.PrintRune(digit)
}
