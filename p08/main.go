package main

//Состояние позволяет динамически изменять поведение объекта при смене его состояния.
//Поведения, зависящие от состояния, переезжают в отдельные типы.
//Первоначальный тип хранит ссылку на один из таких объектов-состояний и делегирует ему работу.

//+
//1) Избавляет от множества больших условных операторов.
//2) Код связанный с одним состоянием находится в одном месте.

//-
//Если состояний мало, то использование паттерна неоправданно усложняет код.

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
