package piscine

func ConvMin(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func ToUpper(s string) string {
	result := []rune{}
	for _, r := range s {
		result = append(result, ConvMin(r))
	}
	return string(result)
}
