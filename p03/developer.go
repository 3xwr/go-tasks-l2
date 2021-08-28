package main

import "fmt"

type Developer struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Developer) fullName() {
	fmt.Println("Developer ", d.FirstName, d.LastName)
}

func (d Developer) accept(v Visitor) {
	v.visitDeveloper(d)
}
