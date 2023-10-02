package piscine

func Split(s, sep string) []string {
	sepLen := len(sep)
	words := make([]string, 0)
	start := 0

	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			if start < i {
				words = append(words, s[start:i])
			}
			i += sepLen - 1
			start = i + 1
		}
	}
	if start < len(s) {
		words = append(words, s[start:])
	}
	return words
}
