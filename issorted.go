package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	f1 := true
	f2 := true
	lens := 0
	for range tab {
		lens++
	}
	for i := 0; i < lens-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			f1 = false
		} else if f(tab[i], tab[i+1]) < 0 {
			f2 = false
		}
	}
	return f1 || f2
}

func f(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
