package main

type orderStrategy interface {
	order(i []item)
}
