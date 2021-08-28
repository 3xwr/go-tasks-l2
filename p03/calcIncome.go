package main

import "fmt"

type calcIncome struct {
	bonusRate int
}

func (c calcIncome) visitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

func (c calcIncome) visitDirector(d Director) {
	fmt.Println(d.Income + d.Income*c.bonusRate/10)
}
