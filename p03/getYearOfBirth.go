package main

import "fmt"

type getYearOfBirth struct {
	yearOfBirth int
}

func (g getYearOfBirth) visitDeveloper(d Developer) {
	fmt.Println(2021 - d.Age)
}

func (g getYearOfBirth) visitDirector(d Director) {
	fmt.Println(2021 - d.Age)
}
