package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for v := range a {
		result[v] = f(a[v])
	}
	return result
}

/*func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}*/
