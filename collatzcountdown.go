package piscine

func CollatzCountdown(start int) int {
	cpt := 0
	if start > 0 {
		for start > 1 {
			cpt++
			if start%2 == 0 {
				start = start / 2
			} else {
				start = 3*start + 1
			}
		}
		return cpt
	} else {
		n := -1
		return n
	}
}
