package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	result := 0
	sign := 1
	signfound := false
	for _, r := range runes {
		if r >= '0' && r <= '9' {
			signfound = true
			result = result*10 + int(r-'0')
		} else if r == '+' && !signfound {
			sign = 1
		} else if r == '-' && !signfound {
			sign = -1
		}
	}
	return result * sign
}
