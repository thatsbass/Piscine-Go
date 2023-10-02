package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	position := 0
	for l != nil {
		if position == pos {
			return l
		}
		l = l.Next
		position++
	}
	return nil
}
