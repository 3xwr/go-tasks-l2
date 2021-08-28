package main

import "fmt"

type lowestpriceStrategy struct {
}

func (l *lowestpriceStrategy) order(i []item) {
	lowestPrice := i[0].price
	bestItem := i[0]
	for _, v := range i {
		if v.price < lowestPrice {
			lowestPrice = v.price
			bestItem = v
		}
	}
	fmt.Println("Ordering item with the lowest price: ", bestItem)
}
