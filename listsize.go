package piscine

func ListSize(l *List) int {
	cpt := 0
	temp := l.Head
	for temp != nil {
		cpt++
		temp = temp.Next
	}
	return cpt
}
