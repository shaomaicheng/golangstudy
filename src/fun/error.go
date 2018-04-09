package main

import (
	"os"
	. "fmt"
	"errors"
	"math"
)

func main() {

	defer func() {
		Println("c")
		if err := recover(); err != nil {
			Println(err)
		}
		Println("d")
	}()
	f()


	Println(sqrt(22.65))
	Println(sqrt(-100))
}

var user = os.Getenv("USER")

func init() {
	Println(user)
	if user == "" {
		panic("no value for $USER")
	}
}

func f() {
	Println("a")
	panic(55)
	Println("b")

	Println("f")
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

type error interface{
	Error() string
}

func sqrt(f float64) (float64, error)  {
	if f < 0 {
		return 0, errors.New("负数平方根不能是0")
	}

	return math.Sqrt(f),  nil
}