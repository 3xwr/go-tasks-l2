package main

import "fmt"

//Строитель отделяет конструирование сложного объекта от его представления так,
//что в результате одного и того же процесса конструирования могут получаться разные представления.

//+
//1) Позволяет строить объекты пошагово с вызовом только нужных шагов
//2) Позволяет создавать разные представления одного объекта (Деревяный и железнобетонный дома)
//3) Изолирует код сборки объекта от бизнес-логики

//-
//1)Код программы усложняется из-за наличия дополнительных типов
//2)Привязанность к конкретным типам строителей
func main() {
	//Строитель с директором
	spicyBuilder := getBuilder("spicy")
	hawaiianBuilder := getBuilder("hawaiian")

	director := newDirector(hawaiianBuilder)
	hawaiianPizza := director.buildPizza()

	fmt.Printf("Hawaiian Pizza Dough Type: %s\n", hawaiianPizza.dough)
	fmt.Printf("Hawaiian Pizza Sauce Type: %s\n", hawaiianPizza.sauce)
	fmt.Printf("Hawaiian Pizza Topping Type: %s\n", hawaiianPizza.topping)

	director.setBuilder(spicyBuilder)
	spicyPizza := director.buildPizza()

	fmt.Printf("Spicy Pizza Dough Type: %s\n", spicyPizza.dough)
	fmt.Printf("Spicy Pizza Sauce Type: %s\n", spicyPizza.sauce)
	fmt.Printf("Spicy Pizza Topping Type: %s\n", spicyPizza.topping)
}
