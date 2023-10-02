package piscine

func RecursiveFactorial(nb int) int {
	limits := 20
	if nb < 0 || nb > limits {
		return 0
	}
	if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
