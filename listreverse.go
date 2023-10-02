package piscine

func ListReverse(l *List) {
	new_link := &List{}
	for l.Head != nil {
		ListPushFront(new_link, l.Head.Data)
		l.Head = l.Head.Next
	}
	*l = *new_link
}
