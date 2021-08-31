package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

//Программа выведет 2 1
//Последний defer выполняяется первым (LIFO порядок).
//defer имеет доступ к именованным возвращаемым аргументам, из-за этого в одной из функций defer
//меняет значение, а в другой - нет.
