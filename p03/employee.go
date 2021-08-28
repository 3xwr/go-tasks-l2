package main

type Employee interface {
	fullName()
	accept(Visitor)
}
