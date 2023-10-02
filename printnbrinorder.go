package piscine

import "github.com/01-edu/z01"

func Tri(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	size := Count(n)
	NewTab := make([]int, size)
	for i := 0; i < size; i++ {
		a := n % 10
		n = n / 10
		NewTab[i] = a
	}
	Tri(NewTab)
	for j := 0; j < size; j++ {
		str := string(rune(NewTab[j]) + '0')
		r := []rune(str)
		z01.PrintRune(r[0])
	}
}

func Count(number int) int {
	counter := 1

	for number > 10 {
		counter++
		number = number / 10
	}
	return counter
}
