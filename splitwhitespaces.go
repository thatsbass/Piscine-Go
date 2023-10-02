package piscine

func Separator(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func SplitWhiteSpaces(s string) []string {
	tab := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if Separator(s[i]) {
			if start < i {
				tab = append(tab, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		tab = append(tab, s[start:])
	}
	return tab
}
