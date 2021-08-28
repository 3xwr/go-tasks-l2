package main

import "fmt"

//Фабричный метод позволяет создавать различные продукты без указания конкретных структур продуктов.

//+
//1) Избавляет тип от привязки к конкретным классам продуктов.
//2) Выделяет код производства продуктов в одно место.
//3) Упрощает добавление новых продуктов в программу.

//-
//Усложнение кода и возможность создания зависимых иерархий

func main() {
	ferrari, err := getCar("ferrari")
	if err == nil {
		printInfo(ferrari)
	}
	lada, err := getCar("lada")
	if err == nil {
		printInfo(lada)
	}
	porsche, err := getCar("porsche")
	if err == nil {
		printInfo(porsche)
	}

}

func printInfo(c icar) {
	fmt.Printf("Car: %s\n", c.getName())
	fmt.Printf("Speed: %d\n", c.getSpeed())
}
