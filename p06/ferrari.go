package main

type ferrari struct {
	car
}

func newFerrari() icar {
	return &ferrari{
		car: car{
			name:  "Ferrari",
			speed: 400,
		},
	}
}
