package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	all_bytes := []byte(s)
	n := 0
	sign := 1
	if len(all_bytes) == 0 {
		return 0
	}
	if all_bytes[0] == byte('-') {
		sign = -1
		all_bytes = all_bytes[1:]
	} else if all_bytes[0] == byte('+') {
		sign = 1
		all_bytes = all_bytes[1:]
	}
	for _, ch := range all_bytes {
		ch -= '0'
		if ch > 9 || ch < 0 {
			return 0
		}
		n = n*10 + int(ch)
	}
	return n * sign
}

func main() {
	args := os.Args
	args = args[1:]
	if len(args) >= 1 {
		if args[0] == "--upper" {
			for _, ch := range args[1:] {
				ch_int := Atoi(ch)
				to_byte := (ch_int + '`') - 32
				reel_rune := rune(to_byte)
				if reel_rune >= 'A' && reel_rune <= 'Z' {
					z01.PrintRune(reel_rune)
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		} else {
			for _, ch := range args[0:] {
				ch_int := Atoi(ch)
				to_byte := (ch_int + '`')
				reel_rune := rune(to_byte)
				if reel_rune >= 'a' && reel_rune <= 'z' {
					z01.PrintRune(reel_rune)
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
