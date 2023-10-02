package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	count := 0
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' || runes[i] >= 'A' && runes[i] <= 'Z' {
			count++
		}
	}
	return count
}
