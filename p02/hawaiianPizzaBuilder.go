package main

type HawaiianPizzaBuilder struct {
	doughType   string
	sauceType   string
	toppingType string
}

func newHawaiianPizzaBuilder() *HawaiianPizzaBuilder {
	return &HawaiianPizzaBuilder{}
}

func (hwpb *HawaiianPizzaBuilder) setDoughType() {
	hwpb.doughType = "Regular"
}

func (hwpb *HawaiianPizzaBuilder) setSauceType() {
	hwpb.sauceType = "Tomato Sauce"
}

func (hwpb *HawaiianPizzaBuilder) setToppingType() {
	hwpb.toppingType = "Ham and Pineapple"
}

func (hwpb *HawaiianPizzaBuilder) getPizza() Pizza {
	return Pizza{
		dough:   hwpb.doughType,
		sauce:   hwpb.sauceType,
		topping: hwpb.toppingType,
	}
}
