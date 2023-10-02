package piscine

func Swap(a *int, b *int) {
	n := *a
	*a = *b
	*b = n
}
