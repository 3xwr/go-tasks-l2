package main

import "fmt"

type material struct {
	next section
}

func (m *material) execute(t *task) {
	if t.materialArrived {
		fmt.Println("Already have materials")
		m.next.execute(t)
		return
	}
	fmt.Println("Collecting materials")
	t.materialArrived = true
	m.next.execute(t)
}

func (m *material) setNext(next section) {
	m.next = next
}
