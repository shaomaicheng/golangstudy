package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct{
	Human
	school string
}

type Employee struct {
	Human
	company string
}


func (h *Human) SayHi() {
	fmt.Printf("hi i am %s you can call me on %s\n", h.name, h.phone)
}

func (s *Student) SayHi() {
	fmt.Printf("i am student %s. my school is %s\n ", s.name, s.school)
	//fmt.Println(s.Human)
}

func main()  {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "golang inc"}

	mark.SayHi()
	sam.SayHi()
}
