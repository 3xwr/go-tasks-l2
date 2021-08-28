package main

type SpicyPizzaBuilder struct {
	doughType   string
	sauceType   string
	toppingType string
}

func newSpicyPizzaBuilder() *SpicyPizzaBuilder {
	return &SpicyPizzaBuilder{}
}

func (spb *SpicyPizzaBuilder) setDoughType() {
	spb.doughType = "Regular"
}

func (spb *SpicyPizzaBuilder) setSauceType() {
	spb.sauceType = "Spicy Sauce"
}

func (spb *SpicyPizzaBuilder) setToppingType() {
	spb.toppingType = "Pepperoni"
}

func (spb *SpicyPizzaBuilder) getPizza() Pizza {
	return Pizza{
		dough:   spb.doughType,
		sauce:   spb.sauceType,
		topping: spb.toppingType,
	}
}
