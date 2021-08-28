package main

type iBuilder interface {
	setDoughType()
	setSauceType()
	setToppingType()
	getPizza() Pizza
}

func getBuilder(pizzaType string) iBuilder {
	if pizzaType == "hawaiian" {
		return &HawaiianPizzaBuilder{}
	}

	if pizzaType == "spicy" {
		return &SpicyPizzaBuilder{}
	}

	return nil
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildPizza() Pizza {
	d.builder.setDoughType()
	d.builder.setSauceType()
	d.builder.setToppingType()
	return d.builder.getPizza()
}
