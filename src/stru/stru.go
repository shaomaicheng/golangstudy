package main

import ."fmt"

func main() {
	var p People
	p.age = 20
	p.name = "p1"

	Println(p)


	var pp *People = &p
	Printf("name: %s; age: %d\n", pp.name, pp.age)

	pp.name = "p2"
	pp.age = 22

	Println(*pp)
	Println(p)  //结构体指针被改变
}

type People struct {
	name string
	age int
}