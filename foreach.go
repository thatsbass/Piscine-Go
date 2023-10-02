package piscine

func ForEach(f func(int), a []int) {
	for v := range a {
		f(a[v])
	}
}
