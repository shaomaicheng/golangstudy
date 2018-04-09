package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	fmt.Println(sum, "写入")
	c <- sum
}

func main() {
	a := []int {7, 2, 8, -9, 4, 0}
	c := make(chan int, 1)


	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)

	x := <- c
	y := <- c

	fmt.Println(x, y)

	ctest := make(chan int, 3)

	go postmsg(ctest)
	//for i := range ctest {
	//	fmt.Println(i)
	//}
	for {
		z,ok := <- ctest
		if !ok {
			fmt.Println("channle 关闭")
			break
		} else {
			fmt.Println(z)
		}
	}
}


func postmsg(ctest chan int) {
	ctest <- 1
	ctest <- 2
	ctest <-3
	close(ctest)
}