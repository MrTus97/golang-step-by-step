package main

import (
	"fmt"
	aa "golang-step-by-step/helper"

	validator "validator"
)

func addItem(list ...int) []int {
	list[0] = 1111
	return list
}

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

	// Array test
	array := [3]int{1, 2, 3}
	array1 := array[1:3]

	array1 = append(array1, 100)
	array2 := append(array1, 200)
	array1[0] = 999

	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println(array2)

	fmt.Println("=========")
	list := addItem(1, 2, 3, 4)
	fmt.Println(list)

	var myMap map[string]int

	if myMap == nil {
		fmt.Println("Mới khởi tạo thì là nil nè")
	}

	// Muốn không nil và dùng được thì phải khởi tạo giá trị khi khai báo
	var myMap2 = map[string]int{
		"k1": 1,
	}
	fmt.Println(myMap2)

	// Muốn tạo thì dùng cái này
	var myMap3 = make(map[string]int)
	myMap3["key1"] = 1

	fmt.Println(myMap3)

	fmt.Println("======== Pointer ==========")
	pointer_value := 100
	var pointer *int
	pointer = &pointer_value
	fmt.Println(pointer)
	fmt.Printf("%T\n", pointer)
	fmt.Println("------------")
	var pointer2 = new(int)
	pointer2 = &pointer_value
	fmt.Println(pointer2)
	fmt.Printf("%T\n", pointer)
	*pointer2 = 999
	fmt.Println(pointer_value)
	fmt.Println("------------")
	pointer_array_value := []int{1, 2, 3}
	var pointer_array *[]int
	pointer_array = &pointer_array_value
	fmt.Println(pointer_array)
	fmt.Println("======== End Pointer ==========")

	struct_demo()
}
