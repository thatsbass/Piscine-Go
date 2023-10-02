package piscine

func BasicAtoi2(s string) int {
	var r int
	lenght := len(s)
	for i := 0; i < lenght; i++ {
		intValue := int(s[i] - '0')
		if s[i] >= '0' && s[i] <= '9' {
			r = r*10 + intValue
		} else {
			return 0
		}
	}
	return r
}
