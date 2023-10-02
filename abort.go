package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	SortIntegerTable(tab)
	return tab[2]
}
