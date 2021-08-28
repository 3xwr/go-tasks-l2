package main

import "fmt"

type app struct {
	data string
}

func (a *app) save() {
	fmt.Println("Saving data: ", a.data)
}
