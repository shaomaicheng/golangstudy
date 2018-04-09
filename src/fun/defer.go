package main

import "fmt"

func main() {

	//for i := 0; i < 5; i++ {
	//	defer fmt.Println("%d", i)
	//}


	slice := []int {1,2,3,4,5,6}
	fmt.Println(filter(slice, isOdd))
}


type testInt func(int) bool  //声明函数类型

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 0 {

		return true
		
	}

	return false
}

func filter(slice []int, f testInt)  []int {

	var res []int

	for _, value := range slice {
		if f(value) {
			res = append(res, value)
		}
	}
	return res
	
}