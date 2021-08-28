package main

//Стратегия позволяет вынести наборы алгоритмов отдельно и делает их взаимозаменимыми.
//Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу.
//Программа может подменить этот объект другим, если требуется иной способ решения задачи.

//+
//1) Легко быстро заменить алгоритм при необходимости
//2) Изоляция кода и данных алогритмов от другого кода

//-
//1) Усложнение кода программы
//2) Для выбора между стратегиями необходимо знать разницу между ними

func main() {
	items := []item{}

	item1 := &item{
		price:        6,
		quality:      2,
		deliverytime: 10,
	}

	item2 := &item{
		price:        40,
		quality:      5,
		deliverytime: 16,
	}

	item3 := &item{
		price:        14,
		quality:      3,
		deliverytime: 1,
	}

	items = append(items, *item1)
	items = append(items, *item2)
	items = append(items, *item3)

	order := &order{
		item:          items,
		orderStrategy: &lowestpriceStrategy{},
	}
	order.order()

	order.setOrderStrategy(&bestqualityStrategy{})
	order.order()

	order.setOrderStrategy(&lowestdeliveryStrategy{})
	order.order()
}
