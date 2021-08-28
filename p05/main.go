package main

//Цепочка обязанностей позволяет передавать запрос по цепочке потенциальных обработчиков, пока один из них
//не выполнит запрос.

//+
//1) Уменьшает зависимость между клиентом и обработчиками.
//2) Реализует принцип единственной обязанности
//3) Реализует принцип закрытости/открытости

//-
//Запрос может остаться необработанным.

func main() {
	material := &material{}
	assembly := &assembly{}
	packaging := &packaging{}
	delivery := &delivery{}

	material.setNext(assembly)
	assembly.setNext(packaging)
	packaging.setNext(delivery)

	task := &task{name: "iron shovel"}
	material.execute(task)
}
