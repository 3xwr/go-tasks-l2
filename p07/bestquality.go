package main

import "fmt"

type bestqualityStrategy struct {
}

func (b *bestqualityStrategy) order(i []item) {
	bestQual := i[0].quality
	bestItem := i[0]
	for _, v := range i {
		if v.quality > bestQual {
			bestQual = v.quality
			bestItem = v
		}
	}
	fmt.Println("Ordering item with the best quality: ", bestItem)
}
