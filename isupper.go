package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 'A' && runes[i] <= 'Z') {
			return false
		}
	}
	return true
}
