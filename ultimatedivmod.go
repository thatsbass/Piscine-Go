package piscine

func UltimateDivMod(a *int, b *int) {
	n := *a / *b
	m := *a % *b
	*a = n
	*b = m
}
