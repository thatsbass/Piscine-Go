package piscine

import (
	"github.com/01-edu/z01"
)

func isSeparator(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, ch := range word {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
