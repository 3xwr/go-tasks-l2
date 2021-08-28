package main

type lada struct {
	car
}

func newLada() icar {
	return &lada{
		car: car{
			name:  "Lada",
			speed: 25,
		},
	}
}
