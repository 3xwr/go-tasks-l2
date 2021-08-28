package main

type porsche struct {
	car
}

func newPorsche() icar {
	return &porsche{
		car: car{
			name:  "Porsche",
			speed: 320,
		},
	}
}
