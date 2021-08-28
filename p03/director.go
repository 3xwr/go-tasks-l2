package main

import "fmt"

type Director struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Director) fullName() {
	fmt.Println("Director ", d.FirstName, d.LastName)
}

func (d Director) accept(v Visitor) {
	v.visitDirector(d)
}
