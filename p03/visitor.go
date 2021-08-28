package main

type Visitor interface {
	visitDeveloper(d Developer)
	visitDirector(d Director)
}
