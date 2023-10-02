package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

type helpMs struct {
	flag        string
	shortenFlag string
	handler     string
}

func obtainValues(value, strsplit string) string {
	values := split(value, strsplit)
	return values[len(values)-1]
}

func split(s, sep string) []string {
	var parts []string
	for {
		i := index(s, sep)
		if i < 0 {
			break
		}
		parts = append(parts, s[:i])
		s = s[i+len(sep):]
	}
	parts = append(parts, s)
	return parts
}

func index(s, substr string) int {
	n := len(s)
	m := len(substr)
	for i := 0; i+m <= n; i++ {
		if s[i:i+m] == substr {
			return i
		}
	}
	return -1
}

func setMs(flag, shortenFlag, handler string) *helpMs {
	helpMs := &helpMs{
		flag:        flag,
		shortenFlag: shortenFlag,
		handler:     handler,
	}
	return helpMs
}

func main() {
	size := len(os.Args)

	if size == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		table := []helpMs{}

		helpMs := setMs("--insert", "-i", "This flag inserts the string into the string passed as argument.")
		table = append(table, *helpMs)
		helpMs = setMs("--order", "-o", "This flag will behave like a boolean, if it is called it will order the argument.")
		table = append(table, *helpMs)

		for _, v := range table {
			fmt.Println(v.flag)
			fmt.Println(" ", v.shortenFlag)
			fmt.Println("	", v.handler)
		}
	} else if size <= 4 {
		var runes []rune
		strToInsert := ""
		var order bool

		for i := 1; i < size; i++ {
			if contains(os.Args[i], "--insert") || contains(os.Args[i], "-i") {
				strToInsert = obtainValues(os.Args[i], "=")
			} else if contains(os.Args[i], "--order") || contains(os.Args[i], "-o") {
				order = true
			} else {
				runes = []rune(os.Args[i])
			}
		}
		if strToInsert != "" {
			concatStr := string(runes) + strToInsert
			runes = []rune(concatStr)
		}
		if order {
			runes = bubbleSort(runes)
		}

		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func contains(s, substr string) bool {
	return index(s, substr) != -1
}

func bubbleSort(arr []rune) []rune {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
