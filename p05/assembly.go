package main

import "fmt"

type assembly struct {
	next section
}

func (a *assembly) execute(t *task) {
	if t.assemblyExecuted {
		fmt.Println("Already assembled")
		a.next.execute(t)
		return
	}
	fmt.Println("Executing assembly")
	t.assemblyExecuted = true
	a.next.execute(t)
}

func (a *assembly) setNext(next section) {
	a.next = next
}
