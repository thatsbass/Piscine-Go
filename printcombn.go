package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	cbn := 1
	var nbNombre string
	if n > 0 || n > 10 {
		for i := 0; i < n; i++ {
			cbn *= 10
			nbNombre += string(rune(i) + '0')
		}
	}
	var r int
	for _, i := range nbNombre {
		num := int(i - '0')
		r = r*10 + num
	}
	for j := 0; j < cbn; j++ {
		rInt := r + j
		if IsPrint(n, rInt) {
			if j != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			cbn2 := 1
			for i := 0; i < n-1; i++ {
				cbn2 *= 10
			}
			if rInt < cbn2 && n != 1 {
				z01.PrintRune('0')
			}
			printNbr(rInt)
		}
	}
	z01.PrintRune('\n')
}

func printNbr(n int) {
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
		printNbr(n / 10)
	}
	runes := '0' + rune(n%10)
	z01.PrintRune(runes)
}

func IsPrint(n, nb int) bool {
	var nb2 int
	bool := true
	p := 10
	for i := 0; i < n; i++ {
		nb2 = nb / 10
		a := nb % 10
		nb = nb2
		if a < p {
			bool = true
			p = a
		} else {
			return false
		}
	}
	return bool
}
