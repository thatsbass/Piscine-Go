package piscine

func BasicAtoi(s string) int {
	var r int
	lenght := len(s)
	for i := 0; i < lenght; i++ {
		intValue := int(s[i] - '0')
		r = r*10 + intValue
	}
	return r
}
