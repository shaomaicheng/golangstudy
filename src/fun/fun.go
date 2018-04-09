package main

import (
	"fmt"
)

func SumAndProduct(A, B int) (add int, Multiplied int) {

	add = A + B
	Multiplied = A * B

	return
	
}


func swipe(a int, b int ) {
	t := a
	a = b
	b = t
}


func rawSWipe(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	x := 3
	y := 4

	xPLUSuy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d", x, y, xPLUSuy)
	fmt.Println()
	fmt.Printf("%d * %d = %d", x, y, xTIMESy)
	fmt.Println()


	swipe(x, y)
	fmt.Println("x: " , x, ";y: ", y)

	rawSWipe(&x, &y)
	fmt.Println("x: " , x, ";y: ", y)
}

