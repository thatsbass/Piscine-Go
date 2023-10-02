package piscine

func Index(s string, toFind string) int {
	R1 := []rune(s)
	R2 := []rune(toFind)
	n1 := len([]rune(s))
	n2 := len([]rune(toFind))
	for i := 0; i < n1; i++ {
		if string(R1[i:i+n2]) == string(R2) {
			return i
		}
	}
	return -1
}
