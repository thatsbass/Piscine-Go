package piscine

func Max(a []int) int {
	max := 0
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}
