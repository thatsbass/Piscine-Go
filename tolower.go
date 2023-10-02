package piscine

func Majconv(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func ToLower(s string) string {
	StockValue := []rune{}
	for _, r := range s {
		StockValue = append(StockValue, Majconv(r))
	}
	return string(StockValue)
}
