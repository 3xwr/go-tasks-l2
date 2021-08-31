package main

import "fmt"

//Программа выведет [77 78 79]
//При объявлении слайса верхний интервал является открытым, т.е. верхняя граница не будет включена в слайс

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}