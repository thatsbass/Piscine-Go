package piscine

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}
	size := max - min
	tab := []int{}

	for i := 0; i < size; i++ {
		tab = append(tab, min+i)
	}
	return tab
}
