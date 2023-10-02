package main

import (
	"fmt"
	"os"
)

func main() {
	param := os.Args[1:]

	for _, i := range param {
		if i == "01" || i == "galaxy" || i == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
