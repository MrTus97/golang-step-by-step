package main

import (
	"fmt"
	aa "golang-step-by-step/helper"

	validator "validator"
)

func main() {
	fmt.Println("Hello world")
	a := 1
	b := 2
	c := a + b
	DI := 3
	fmt.Println(c + DI)
	fmt.Println(aa.SumFromZero(4))

	validator.CheckValidEmail()

	number := 1
	switch number {
	case 1:
		goto handle
		fmt.Println(1)
	handle:
		fmt.Println(2)
		// fallthrough
	case 2:
		fmt.Println(3)
	}
}
