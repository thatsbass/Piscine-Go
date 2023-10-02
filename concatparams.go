package piscine

func ConcatParams(args []string) string {
	size := len(args)
	if size == 0 {
		return ""
	}
	value := args[0]
	for i := 1; i < size; i++ {
		value = Concat(value, "\n"+args[i])
	}
	return value
}
