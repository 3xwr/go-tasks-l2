package main

//Посетитель позволяет добавить поведение в структуру без ее изменения.

//+
//1) Упрощает добавление операций, работающих со сложными структурами объектов.
//2) Объединяет схожие операции в одном месте.
//3) Посетитель может накапливать состояние при обходе структуры элементов (например, дерева).

//-
//1) Паттерн не имеет смысла, если иерархия элементов часто меняется.

func main() {
	dev := Developer{"Ivan", "Ivanov", 1000, 26}
	boss := Director{"Anton", "Antonov", 2500, 38}

	dev.fullName()
	dev.accept(calcIncome{20})
	dev.accept(getYearOfBirth{})

	boss.fullName()
	boss.accept(calcIncome{10})
	boss.accept(getYearOfBirth{})
}
