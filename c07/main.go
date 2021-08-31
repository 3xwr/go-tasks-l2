package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			if a == nil && b == nil {
				close(c)
				break
			}
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					break
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					break
				}
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}

}

//Исходная программа выведет в смешанном порядке числа от 1 до 8, а затем будет выводить нули.
//Каналы на самом деле посылают сообщения формата (msg, ok), где ok = false, а msg = дефолтное значение типа
// в случае закрытия канала.
//Чтение из закрытого канала не является блокирующей операцией, мы не получаем deadlock в merge и просто продолжаем
//читать сообщения о том, что канал закрыт, отсюда и берутся нули.
