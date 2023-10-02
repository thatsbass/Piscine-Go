package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	runes := []rune(s)
	for _, i := range runes {
		z01.PrintRune(i)
	}
}
