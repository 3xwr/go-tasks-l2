package main

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	fmt.Println("Using button to perform command.")
	b.command.execute()
}
