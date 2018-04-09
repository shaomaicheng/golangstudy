package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name string
	age int
}

type Student struct {
	Human
	school string
}

func (h Human) SayHi() {
	fmt.Printf("name: %s, age: %d\n", h.name, h.age)
}

func (h Human) Sing() {
	fmt.Println("sing")
}

type Men interface{
	SayHi()
	Sing()
}

func main()  {
	h := Human{"h1", 20}
	var m Men
	m = h
	m.SayHi()
	m.Sing()

	fmt.Println(reflect.TypeOf(h))

	s := Student{Human{"h2", 20}, "mit"}
	m = s
	m.SayHi()
	m.Sing()

	t := reflect.TypeOf(m)
	fmt.Println(t)
}