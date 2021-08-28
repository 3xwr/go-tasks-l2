package main

import "fmt"

type packaging struct {
	next section
}

func (p *packaging) execute(t *task) {
	if t.packagingExecuted {
		fmt.Println("Already packaged")
		p.next.execute(t)
		return
	}
	fmt.Println("Executing packaging")
	t.packagingExecuted = true
	p.next.execute(t)
}

func (p *packaging) setNext(next section) {
	p.next = next
}
