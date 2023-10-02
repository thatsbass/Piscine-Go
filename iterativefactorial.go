package piscine

func IterativeFactorial(nb int) int {
	limits := 20
	if nb < 0 || nb > limits {
		return 0
	}
	if nb == 0 {
		return 1
	}
	a := 1
	if nb >= 1 {
		for i := 1; i <= nb; i++ {
			a *= i
		}
	}
	return a
}
