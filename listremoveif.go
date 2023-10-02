package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	new_list := &List{}
	for n != nil {
		if !CompStr(data_ref, n.Data) {
			ListPushBack(new_list, n.Data)
		}
		n = n.Next
	}
	*l = *new_list
}
