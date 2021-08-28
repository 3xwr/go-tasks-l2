package main

import "fmt"

type lowestdeliveryStrategy struct {
}

func (l *lowestdeliveryStrategy) order(i []item) {
	lowestTime := i[0].deliverytime
	bestItem := i[0]
	for _, v := range i {
		if v.deliverytime < lowestTime {
			lowestTime = v.deliverytime
			bestItem = v
		}
	}
	fmt.Println("Ordering item with the lowest delivery time: ", bestItem)
}
