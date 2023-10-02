package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	data, _ := ioutil.ReadFile(os.Args[1])
	fmt.Print(string(data))
}
