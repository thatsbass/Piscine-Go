package main

import (
	"os"

	"github.com/01-edu/z01"
)

func RangeSort(arguments []string) {
	n := len(arguments)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if compareStrings(arguments[j], arguments[j+1]) > 0 {
				arguments[j], arguments[j+1] = arguments[j+1], arguments[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func compareStrings(a, b string) int {
	i := 0
	for i < len(a) && i < len(b) {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
		i++
	}

	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}

func printArgumentsInASCIIOrder(arguments []string) {
	RangeSort(arguments)
	for _, arg := range arguments {
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	arguments := os.Args[1:]
	printArgumentsInASCIIOrder(arguments)
}
