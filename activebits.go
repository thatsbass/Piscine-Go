package piscine

func ActiveBits(n int) int {
	active := 0
	for n != 0 {
		if n%2 == 1 {
			active++
		}
		n = n / 2
	}
	return active
}
