package piscine

func CountIf(f func(string) bool, tab []string) int {
	cpt := 0
	for _, v := range tab {
		if f(v) {
			cpt++
		}
	}
	return cpt
}
