package piscine

func ListClear(l *List) {
	pre := l.Head
	temp := l.Head
	if l.Head == nil {
		return
	}
	if temp.Next != nil {
		pre = temp
		temp = temp.Next
	}
	l.Tail = pre
	l.Tail.Next = l.Head
	l.Head = nil
}
