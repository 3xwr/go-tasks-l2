package main

type icar interface {
	setName(name string)
	setSpeed(speed int)
	getName() string
	getSpeed() int
}

type car struct {
	name  string
	speed int
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) setSpeed(speed int) {
	c.speed = speed
}

func (c *car) getSpeed() int {
	return c.speed
}
