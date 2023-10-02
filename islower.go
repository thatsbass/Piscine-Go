package piscine

func IsLower(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 'a' && runes[i] <= 'z') {
			return false
		}
	}
	return true
}
