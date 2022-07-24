package main

import "fmt"

type PeopleInfo struct {
	class string
	email string
	age   string
}

type People struct {
	name string
	info PeopleInfo
}

type Student struct {
	id   int
	name string
}

func struct_demo() {
	//named
	st1 := Student{
		id:   123,
		name: "Ryan",
	}

	fmt.Println(st1)
	fmt.Println(st1.id)

	st2 := Student{1, "TuPT"}
	fmt.Println(st2)

	var st3 Student = struct {
		id   int
		name string
	}{
		id:   777,
		name: "DuyenHTM",
	}

	fmt.Println(st3)

	// anonymous struct
	var anonymous = struct {
		email string
		age   int
	}{
		"ryan@gmail.com",
		27,
	}
	fmt.Println(anonymous)

	// Pointer trỏ tới struct
	st4 := &Student{
		999, "Robin",
	}
	fmt.Println(&st4)
	fmt.Println((*st4).name)
	fmt.Println(st4.name)

	// anonymous fields
	type NoName struct {
		string
		int
	}

	n := NoName{"Nguyen Van A", 56}
	fmt.Println(n)

	// Struct long struct - nested struct
	people1 := People{
		name: "ahihi",
		info: PeopleInfo{
			class: "1",
			email: "2",
			age:   "3",
		},
	}

	fmt.Println(people1)

	people2 := People{
		name: "ahihi123",
	}
	fmt.Println(people2)

	// compare 2 struct
	type structCompare struct {
		id   int
		name string
		// info map[int]int => Nếu có chứa dữ liệu không so sánh đưuọc thì sẽ không so sánh được
	}

	s1 := structCompare{2, "2"}
	s2 := structCompare{1, "2"}
	if s1 == s2 {
		fmt.Println("S1 == s2")
	} else {
		fmt.Println("S1 != s2")
	}

	// zero value

	var st5 Student
	st5.name = "TuPT"
	fmt.Println(st5)
}
