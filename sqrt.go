package piscine

func Sqrt(nb int) int {
	x := 1.0
	N := float64(nb)
	for i := 0; i < 10; i++ {
		x -= (x*x - N) / (2 * x)
	}
	if int(x)*int(x) == int(N) {
		return int(x)
	}
	return 0
}
