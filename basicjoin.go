package piscine

func BasicJoin(elems []string) string {
	size := len(elems)
	NewTab := []rune{}
	for i := 0; i < size; i++ {
		r := elems[i]
		ConvTab := []rune(r)
		for j := 0; j < len(r); j++ {
			NewTab = append(NewTab, ConvTab[j])
		}
	}
	return string(NewTab)
}
