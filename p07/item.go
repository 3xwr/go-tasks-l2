package main

type item struct {
	price        int
	quality      int
	deliverytime int
}

type order struct {
	item          []item
	orderStrategy orderStrategy
}

func (o *order) setOrderStrategy(os orderStrategy) {
	o.orderStrategy = os
}

func (o *order) order() {
	o.orderStrategy.order(o.item)
}
