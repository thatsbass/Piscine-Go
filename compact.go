package piscine

func Compact(ptr *[]string) int {
	var NewtabX []string
	cpt := 0
	for _, i := range *ptr {
		if i != "" {
			cpt++
			NewtabX = append(NewtabX, i)
		}
	}
	*ptr = NewtabX
	return cpt
}
