package piscine

func Join(strs []string, sep string) string {
	sepRune := []rune(sep)
	size := len(strs)
	NewTab := []rune{}
	for i := 0; i < size; i++ {
		r := strs[i]
		ConvTab := []rune(r)
		for j := 0; j < len(r); j++ {
			NewTab = append(NewTab, ConvTab[j])
		}
		if i < size-1 {
			NewTab = append(NewTab, sepRune...)
		}
	}
	return string(NewTab)
}
