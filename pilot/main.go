package main

import (
	"fmt"
)

const AIRCRAFT1 = 1

func main() {
	type Pilot struct {
		Name     string
		Life     float64
		Age      int
		Aircraft int
	}
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}
