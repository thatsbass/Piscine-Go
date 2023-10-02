package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	pos := len([]rune(s))
	if n > pos || n <= 0 {
		return rune(0)
	}
	if n == 1 {
		return runes[0]
	}
	if n != 0 {
		return runes[n-1]
	}
	return runes[n]
}
