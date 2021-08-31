package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

//Программа выведет <nil> false
//Интерфейсы внутри Go выглядят как пара (value, type)
//Foo() возвращает [nil, *os.PathError], что мы сравниваем с [nil,nil] и закономерно получаем false.
//В пустых интерфейсах изначально оба поля имеют значение nil.
//В обычных интерфейсах одно поле имеет какое-то значение, другое поле - значение, особое для типа.
