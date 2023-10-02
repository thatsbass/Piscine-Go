package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= 'A' && runes[i] <= 'Z' || runes[i] >= '0' && runes[i] <= '9') {
			return false
		}
	}
	return true
}
