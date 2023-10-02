package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	tab := make([]int, size)

	for i := 0; i < size; i++ {
		tab[i] = min + i
	}
	return tab
}
