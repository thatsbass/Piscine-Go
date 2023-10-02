package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	runes := []rune(s)
	for _, i := range runes {
		z01.PrintRune(i)
	}
}

func main() {
	if len(os.Args) == 1 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			PrintStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		PrintStr(string(bytes))
		return
	}
	for _, v := range os.Args[1:] {
		file, err := os.Open(v)
		if err != nil {
			PrintStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			PrintStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		PrintStr(string(bytes))
	}
}
