package piscine

func Atoi(s string) int {
	result, sign := 0, 1
	for i, val := range s {
		if i == 0 && (val == '-' || val == '+') {
			if val == '-' {
				sign = -1
			}
			continue
		}
		if val < '0' || val > '9' {
			return 0
		}
		result = result*10 + int(val-'0')
	}
	return result * sign
}
