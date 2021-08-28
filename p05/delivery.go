package main

import "fmt"

type delivery struct {
	next section
}

func (d *delivery) execute(t *task) {
	if t.deliveryExecuted {
		fmt.Println("Already delivered")
		d.next.execute(t)
		return
	}
	fmt.Println("Executing delivery")
	t.deliveryExecuted = true
}

func (d *delivery) setNext(next section) {
	d.next = next
}
